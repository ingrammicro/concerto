package bootstrapping

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/allan-simon/go-singleinstance"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

const (
	//DefaultTimingInterval Default period for looping
	DefaultTimingInterval       = 600 // 600 seconds = 10 minutes
	DefaultTimingSplay          = 360 // seconds
	DefaultThresholdLines       = 10
	DefaultApplyAfterIterations = 4 // iterations
	ProcessLockFile             = "cio-bootstrapping.lock"
	RetriesNumber               = 5
)

type bootstrappingProcess struct {
	startedAt                    time.Time
	finishedAt                   time.Time
	policyfiles                  []policyfile
	attributes                   attributes
	thresholdLines               int
	directoryPath                string
	appliedPolicyfileRevisionIDs map[string]string
}
type attributes struct {
	revisionID string
	rawData    map[string]interface{}
}

type policyfile types.BootstrappingPolicyfile

func (pf policyfile) Name() string {
	return strings.Join([]string{pf.ID, "-", pf.RevisionID}, "")
}

func (pf *policyfile) FileName() string {
	return strings.Join([]string{pf.Name(), "tgz"}, ".")
}

func (pf *policyfile) QueryURL() (string, error) {
	if pf.DownloadURL == "" {
		return "", fmt.Errorf("obtaining URL query: empty download URL")
	}
	url, err := url.Parse(pf.DownloadURL)
	if err != nil {
		return "", fmt.Errorf("parsing URL to extract query: %v", err)
	}
	return fmt.Sprintf("%s?%s", url.Path, url.RawQuery), nil
}

func (pf *policyfile) TarballPath(dir string) string {
	return filepath.Join(dir, pf.FileName())
}

func (pf *policyfile) Path(dir string) string {
	return filepath.Join(dir, pf.Name())
}

func (a *attributes) FileName() string {
	return fmt.Sprintf("attrs-%s.json", a.revisionID)
}

func (a *attributes) FilePath(dir string) string {
	return filepath.Join(dir, a.FileName())
}

// Handle signals
func handleSysSignals(cancelFunc context.CancelFunc) {
	log.Debug("handleSysSignals")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Debug("Ending, signal detected:", <-gracefulStop)
	cancelFunc()
}

// Returns the full path to the tmp directory joined with pid management file name
func lockFilePath() string {
	return filepath.Join(os.TempDir(), ProcessLockFile)
}

func workspaceDir() string {
	return filepath.Join(os.TempDir(), "cio")
}

// Returns the full path to the tmp directory
func generateWorkspaceDir() error {
	dir := workspaceDir()
	dirInfo, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			return err
		}
	} else {
		if !dirInfo.Mode().IsDir() {
			return fmt.Errorf("%s exists but is not a directory", dir)
		}
	}
	return nil
}

// Start the bootstrapping process
func start(c *cli.Context) error {
	log.Debug("start")

	err := generateWorkspaceDir()
	if err != nil {
		return err
	}
	lockFile, err := singleinstance.CreateLockFile(lockFilePath())
	if err != nil {
		return err
	}
	defer lockFile.Close()

	formatter := format.GetFormatter()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go handleSysSignals(cancel)

	config, err := utils.GetConcertoConfig()
	if err != nil {
		formatter.PrintFatal("Couldn't wire up config", err)
	}

	interval := config.BootstrapConfig.IntervalSeconds
	if !(interval > 0) {
		interval = DefaultTimingInterval
	}

	splay := config.BootstrapConfig.SplaySeconds
	if !(splay > 0) {
		splay = DefaultTimingSplay
	}

	applyAfterIterations := config.BootstrapConfig.ApplyAfterIterations
	if !(applyAfterIterations > 0) {
		applyAfterIterations = DefaultApplyAfterIterations
	}

	thresholdLines := c.Int("lines")
	if !(thresholdLines > 0) {
		thresholdLines = DefaultThresholdLines
	}
	log.Debug("routine lines threshold: ", thresholdLines)
	bootstrappingSvc, formatter := cmd.WireUpBootstrapping(c)

	if config.BootstrapConfig.RunOnce {
		return runBootstrapOnce(ctx, bootstrappingSvc, formatter, thresholdLines, interval, splay)
	}
	return runBootstrapPeriodically(ctx, bootstrappingSvc, formatter, applyAfterIterations, thresholdLines, interval, splay)
}

// Stop the bootstrapping process
func stop(c *cli.Context) error {
	log.Debug("cmdStop")

	formatter := format.GetFormatter()
	if err := utils.StopProcess(lockFilePath()); err != nil {
		formatter.PrintFatal("cannot stop the bootstrapping process", err)
	}

	log.Info("Bootstrapping routine successfully stopped")
	return nil
}

func runBootstrapPeriodically(ctx context.Context, bootstrappingSvc *blueprint.BootstrappingService, formatter format.Formatter, applyAfterIterations, thresholdLines, interval, splay int) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var blueprintConfig *types.BootstrappingConfiguration
	var noPolicyfileApplicationIterations int
	var lastPolicyfileApplicationErr, err error
	for {
		var updated bool
		blueprintConfig, updated, err = getBlueprintConfig(ctx, bootstrappingSvc, blueprintConfig, formatter)
		if err == nil {
			if updated || lastPolicyfileApplicationErr != nil || noPolicyfileApplicationIterations >= applyAfterIterations {
				noPolicyfileApplicationIterations = -1
				lastPolicyfileApplicationErr = applyPolicyfiles(ctx, bootstrappingSvc, blueprintConfig, formatter, thresholdLines)
			}
		}
		noPolicyfileApplicationIterations++

		// Sleep for a configured amount of time plus a random amount of time (10 minutes plus 0 to 5 minutes, for instance)
		ticker := time.NewTicker(time.Duration(interval+r.Intn(int(splay))) * time.Second)

		select {
		case <-ticker.C:
			log.Debug("ticker")
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing bootstrapping")
		}
		ticker.Stop()
		if ctx.Err() != nil {
			break
		}
	}
	return nil
}

func runBootstrapOnce(ctx context.Context, bootstrappingSvc *blueprint.BootstrappingService, formatter format.Formatter, thresholdLines, interval, splay int) error {
	blueprintConfig, _, err := getBlueprintConfig(ctx, bootstrappingSvc, nil, formatter)
	if err == nil {
		err = applyPolicyfiles(ctx, bootstrappingSvc, blueprintConfig, formatter, thresholdLines)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; err != nil && i < 3; i++ {
		// Sleep for a configured amount of time plus a random amount of time (10 minutes plus 0 to 5 minutes, for instance)
		ticker := time.NewTicker(time.Duration(interval+r.Intn(int(splay))) * time.Second)
		select {
		case <-ticker.C:
			log.Debug("ticker")
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("interrupting bootstrapping")
			break
		}
		ticker.Stop()
		blueprintConfig, _, err = getBlueprintConfig(ctx, bootstrappingSvc, nil, formatter)
		if err == nil {
			err = applyPolicyfiles(ctx, bootstrappingSvc, blueprintConfig, formatter, thresholdLines)
		}
	}
	return err
}

func getBlueprintConfig(ctx context.Context, bootstrappingSvc *blueprint.BootstrappingService, previousBlueprintConfig *types.BootstrappingConfiguration, formatter format.Formatter) (*types.BootstrappingConfiguration, bool, error) {
	log.Debug("getBlueprintConfig")
	// Inquire about desired configuration changes to be applied by querying the `GET /blueprint/configuration` endpoint. This will provide a JSON response with the desired configuration changes
	blueprintConfig, status, err := bootstrappingSvc.GetBootstrappingConfiguration()
	if err == nil && status != 200 {
		err = fmt.Errorf("received non-ok %d response", status)
	}
	if err != nil {
		formatter.PrintError("couldn't receive bootstrapping data", err)
		return previousBlueprintConfig, false, err
	}
	updated := previousBlueprintConfig == nil
	if !updated {
		updated = previousBlueprintConfig.AttributeRevisionID != blueprintConfig.AttributeRevisionID
		updated = updated || len(previousBlueprintConfig.Policyfiles) != len(blueprintConfig.Policyfiles)
	}
	if !updated {
		for i, cp := range blueprintConfig.Policyfiles {
			pp := previousBlueprintConfig.Policyfiles[i]
			updated = cp.ID != pp.ID || cp.RevisionID != pp.RevisionID
			if updated {
				break
			}
		}
	}
	return blueprintConfig, updated, ctx.Err()
}

// Subsidiary routine for commands processing
func applyPolicyfiles(ctx context.Context, bootstrappingSvc *blueprint.BootstrappingService, blueprintConfig *types.BootstrappingConfiguration, formatter format.Formatter, thresholdLines int) error {
	log.Debug("applyPolicyfiles")
	err := generateWorkspaceDir()
	if err != nil {
		formatter.PrintError("couldn't generated workspace directory", err)
		return err
	}
	bsProcess := &bootstrappingProcess{
		startedAt:                    time.Now().UTC(),
		thresholdLines:               thresholdLines,
		directoryPath:                workspaceDir(),
		appliedPolicyfileRevisionIDs: make(map[string]string),
	}

	// proto structures
	err = initializePrototype(blueprintConfig, bsProcess)
	if err != nil {
		formatter.PrintError("couldn't initialize prototype", err)
		return err
	}
	// For every policyfile, ensure its tarball (downloadable through their download_url) has been downloaded to the server ...
	err = downloadPolicyfiles(ctx, bootstrappingSvc, bsProcess)
	if err != nil {
		formatter.PrintError("couldn't download policy files", err)
		return err
	}
	//... and clean off any tarball that is no longer needed.
	err = cleanObsoletePolicyfiles(bsProcess)
	if err != nil {
		formatter.PrintError("couldn't clean obsolete policy files", err)
		return err
	}
	// Store the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
	err = saveAttributes(bsProcess)
	if err != nil {
		formatter.PrintError("couldn't save attributes for policy files", err)
		return err
	}
	// Process tarballs policies
	err = processPolicyfiles(bootstrappingSvc, bsProcess)
	// Finishing time
	bsProcess.finishedAt = time.Now().UTC()

	// Inform the platform of applied changes via a `PUT /blueprint/applied_configuration` request with a JSON payload similar to
	log.Debug("reporting applied policy files")
	reportErr := reportAppliedConfiguration(bootstrappingSvc, bsProcess)
	if reportErr != nil {
		formatter.PrintError("couldn't report applied status for policy files", err)
		return err
	}
	return err
}

func initializePrototype(bsConfiguration *types.BootstrappingConfiguration, bsProcess *bootstrappingProcess) error {
	log.Debug("initializePrototype")

	// Attributes
	bsProcess.attributes.revisionID = bsConfiguration.AttributeRevisionID
	bsProcess.attributes.rawData = bsConfiguration.Attributes

	// Policies
	for _, bsConfPolicyfile := range bsConfiguration.Policyfiles {
		bsProcess.policyfiles = append(bsProcess.policyfiles, policyfile(bsConfPolicyfile))
	}
	log.Debug(bsProcess)
	return nil
}

// downloadPolicyfiles For every policy file, ensure its tarball (downloadable through their download_url) has been downloaded to the server ...
func downloadPolicyfiles(ctx context.Context, bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) error {
	log.Debug("downloadPolicyfiles")

	for _, bsPolicyfile := range bsProcess.policyfiles {
		tarballPath := bsPolicyfile.TarballPath(bsProcess.directoryPath)
		log.Debug("downloading: ", tarballPath)
		queryURL, err := bsPolicyfile.QueryURL()
		if err != nil {
			return err
		}
		_, status, err := bootstrappingSvc.DownloadPolicyfile(queryURL, tarballPath)
		if err == nil && status != 200 {
			err = fmt.Errorf("obtained non-ok response when downloading policyfile %s", queryURL)
		}
		if err != nil {
			return err
		}
		if err = utils.Untar(ctx, tarballPath, bsPolicyfile.Path(bsProcess.directoryPath)); err != nil {
			return err
		}
	}
	return nil
}

// cleanObsoletePolicyfiles cleans off any tarball that is no longer needed.
func cleanObsoletePolicyfiles(bsProcess *bootstrappingProcess) error {
	log.Debug("cleanObsoletePolicyfiles")

	// evaluates working folder
	deletableFiles, err := ioutil.ReadDir(bsProcess.directoryPath)
	if err != nil {
		return err
	}

	// builds an array of currently processable files at this looping time
	currentlyProcessableFiles := []string{bsProcess.attributes.FileName()} // saved attributes file name
	for _, bsPolicyFile := range bsProcess.policyfiles {
		currentlyProcessableFiles = append(currentlyProcessableFiles, bsPolicyFile.FileName()) // Downloaded tgz file names
		currentlyProcessableFiles = append(currentlyProcessableFiles, bsPolicyFile.Name())     // Uncompressed folder names
	}

	// removes from deletableFiles array the policy files currently applied
	for _, f := range deletableFiles {
		if !utils.Contains(currentlyProcessableFiles, f.Name()) {
			log.Debug("removing: ", f.Name())
			if err := os.RemoveAll(filepath.Join(bsProcess.directoryPath, f.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}

// saveAttributes stores the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
func saveAttributes(bsProcess *bootstrappingProcess) error {
	log.Debug("saveAttributes")

	attrs, err := json.Marshal(bsProcess.attributes.rawData)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(bsProcess.attributes.FilePath(bsProcess.directoryPath), attrs, 0600); err != nil {
		return err
	}
	return nil
}

// processPolicyfiles applies for each policy the required chef commands, reporting in bunches of N lines
func processPolicyfiles(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) error {
	log.Debug("processPolicyfiles")

	for _, bsPolicyfile := range bsProcess.policyfiles {
		command := fmt.Sprintf("chef-client -z -j %s", bsProcess.attributes.FilePath(bsProcess.directoryPath))
		policyfileDir := bsPolicyfile.Path(bsProcess.directoryPath)
		var renamedPolicyfileDir string
		if runtime.GOOS == "windows" {
			renamedPolicyfileDir = policyfileDir
			policyfileDir = filepath.Join(bsProcess.directoryPath, "active")
			err := os.Rename(renamedPolicyfileDir, policyfileDir)
			if err != nil {
				return fmt.Errorf("could not rename %s as %s: %v", renamedPolicyfileDir, policyfileDir, err)
			}
			command = fmt.Sprintf("SET \"PATH=%%PATH%%;C:\\ruby\\bin;C:\\opscode\\chef\\bin;C:\\opscode\\chef\\embedded\\bin\"\n%s", command)
		}
		command = fmt.Sprintf("cd %s\n%s", policyfileDir, command)

		log.Debug(command)

		// Custom method for chunks processing
		fn := func(chunk string) error {
			log.Debug("sendChunks")
			err := utils.Retry(RetriesNumber, time.Second, func() error {
				log.Debug("Sending: ", chunk)

				commandIn := map[string]interface{}{
					"stdout": chunk,
				}

				_, statusCode, err := bootstrappingSvc.ReportBootstrappingLog(&commandIn)
				switch {
				// 0<100 error cases??
				case statusCode == 0:
					return fmt.Errorf("communication error %v %v", statusCode, err)
				case statusCode >= 500:
					return fmt.Errorf("server error %v %v", statusCode, err)
				case statusCode >= 400:
					return fmt.Errorf("client error %v %v", statusCode, err)
				default:
					return nil
				}
			})

			if err != nil {
				return fmt.Errorf("cannot send the chunk data, %v", err)
			}
			return nil
		}

		exitCode, err := utils.RunContinuousCmd(fn, command, -1, bsProcess.thresholdLines)
		if err == nil && exitCode != 0 {
			err = fmt.Errorf("policyfile application exited with %d code", exitCode)
		}
		if err != nil {
			return err
		}

		log.Info("completed: ", exitCode)
		bsProcess.appliedPolicyfileRevisionIDs[bsPolicyfile.ID] = bsPolicyfile.RevisionID
		if renamedPolicyfileDir != "" {
			err = os.Rename(policyfileDir, renamedPolicyfileDir)
			if err != nil {
				return fmt.Errorf("could not rename %s as %s back: %v", policyfileDir, renamedPolicyfileDir, err)
			}
		}
	}
	return nil
}

// reportAppliedConfiguration Inform the platform of applied changes
func reportAppliedConfiguration(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) error {
	log.Debug("reportAppliedConfiguration")

	payload := map[string]interface{}{
		"started_at":              bsProcess.startedAt,
		"finished_at":             bsProcess.finishedAt,
		"policyfile_revision_ids": bsProcess.appliedPolicyfileRevisionIDs,
		"attribute_revision_id":   bsProcess.attributes.revisionID,
	}
	return bootstrappingSvc.ReportBootstrappingAppliedConfiguration(&payload)
}
