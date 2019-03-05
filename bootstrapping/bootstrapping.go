package bootstrapping

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"math/rand"
	"fmt"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

const (
	//DefaultTimingInterval Default period for looping
	DefaultTimingInterval = 600 // 600 seconds = 10 minutes
	DefaultTimingSplay    = 360 // seconds
	DefaultThresholdLines = 10
	ProcessIDFile         = "imco-bootstrapping.pid"
	RetriesNumber         = 5
)

type bootstrappingProcess struct {
	startedAt      string
	finishedAt     string
	policyFiles    []*policyFile
	attributes     attributes
	thresholdLines int
	directoryPath  string
}
type attributes struct {
	revisionID string
	fileName   string
	filePath   string
	rawData    *json.RawMessage
}

type policyFile struct {
	id          string
	revisionID  string
	name        string
	fileName    string
	tarballURL  string
	queryURL    string
	tarballPath string
	folderPath  string

	downloaded   bool
	uncompressed bool
	executed     bool
}

var allPolicyFilesSuccessfullyApplied bool

// Handle signals
func handleSysSignals(cancelFunc context.CancelFunc) {
	log.Debug("handleSysSignals")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Debug("Ending, signal detected:", <-gracefulStop)
	cancelFunc()
}

// Returns the full path to the tmp folder joined with pid management file name
func getProcessIDFilePath() string {
	return strings.Join([]string{os.TempDir(), string(os.PathSeparator), ProcessIDFile}, "")
}

// Returns the full path to the tmp folder
func getProcessingFolderFilePath() string {
	dir := strings.Join([]string{os.TempDir(), string(os.PathSeparator), "imco", string(os.PathSeparator)}, "")
	os.Mkdir(dir, 0777)
	return dir
}

// Start the bootstrapping process
func start(c *cli.Context) error {
	log.Debug("start")

	formatter := format.GetFormatter()
	if err := utils.SetProcessIdToFile(getProcessIDFilePath()); err != nil {
		formatter.PrintFatal("cannot create the pid file", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSysSignals(cancel)

	bootstrappingRoutine(ctx, c)

	return nil
}

// Stop the bootstrapping process
func stop(c *cli.Context) error {
	log.Debug("cmdStop")

	formatter := format.GetFormatter()
	if err := utils.StopProcess(getProcessIDFilePath()); err != nil {
		formatter.PrintFatal("cannot stop the bootstrapping process", err)
	}

	log.Info("Bootstrapping routine successfully stopped")
	return nil
}

// Main bootstrapping background routine
func bootstrappingRoutine(ctx context.Context, c *cli.Context) {
	log.Debug("bootstrappingRoutine")

	timingInterval := c.Int64("interval")
	if !(timingInterval > 0) {
		timingInterval = DefaultTimingInterval
	}

	timingSplay := c.Int64("splay")
	if !(timingSplay > 0) {
		timingSplay = DefaultTimingSplay
	}

	thresholdLines := c.Int("lines")
	if !(thresholdLines > 0) {
		thresholdLines = DefaultThresholdLines
	}
	log.Debug("routine lines threshold: ", thresholdLines)

	bootstrappingSvc, formatter := cmd.WireUpBootstrapping(c)
	for {
		applyPolicyfiles(bootstrappingSvc, formatter, thresholdLines)

		// Sleep for a configured amount of time plus a random amount of time (10 minutes plus 0 to 5 minutes, for instance)
		ticker := time.NewTicker(time.Duration(timingInterval + int64(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(int(timingSplay)))) * time.Second)

		select {
		case <- ticker.C:
			log.Debug("ticker")
		case <- ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing bootstrapping")
		}
		ticker.Stop()
		if ctx.Err() != nil {
			break
		}
	}
}

// Subsidiary routine for commands processing
func applyPolicyfiles(bootstrappingSvc *blueprint.BootstrappingService, formatter format.Formatter, thresholdLines int) {
	log.Debug("applyPolicyfiles")

	// Inquire about desired configuration changes to be applied by querying the `GET /blueprint/configuration` endpoint. This will provide a JSON response with the desired configuration changes
	bsConfiguration, status, err := bootstrappingSvc.GetBootstrappingConfiguration()
	if err != nil {
		formatter.PrintError("couldn't receive bootstrapping data", err)
	} else {
		if status == 200 {
			bsProcess := new(bootstrappingProcess)
			// Starting time
			bsProcess.startedAt = time.Now().UTC().String()
			bsProcess.thresholdLines = thresholdLines
			bsProcess.directoryPath = getProcessingFolderFilePath()

			// proto structures
			initializePrototype(bsConfiguration, bsProcess)

			// For every policyFile, ensure its tarball (downloadable through their download_url) has been downloaded to the server ...
			downloadPolicyFiles(bootstrappingSvc, bsProcess)

			//... and clean off any tarball that is no longer needed.
			cleanObsoletePolicyFiles(bsProcess)

			// Store the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
			saveAttributes(bsProcess)

			// Process tarballs policies
			processPolicyFiles(bootstrappingSvc, bsProcess)

			// Finishing time
			bsProcess.finishedAt = time.Now().UTC().String()

			// Inform the platform of applied changes via a `PUT /blueprint/applied_configuration` request with a JSON payload similar to
			log.Debug("reporting applied policy files")
			reportAppliedConfiguration(bootstrappingSvc, bsProcess)

			completeBootstrappingSequence(bsProcess)
		}
	}
}

func initializePrototype(bsConfiguration *types.BootstrappingConfiguration, bsProcess *bootstrappingProcess) {
	log.Debug("initializePrototype")

	// Attributes
	bsProcess.attributes.revisionID = bsConfiguration.AttributeRevisionID
	bsProcess.attributes.fileName = strings.Join([]string{"attrs-", bsProcess.attributes.revisionID, ".json"}, "")
	bsProcess.attributes.filePath = strings.Join([]string{bsProcess.directoryPath, bsProcess.attributes.fileName}, "")
	bsProcess.attributes.rawData = bsConfiguration.Attributes

	// Policies
	for _, bsConfPolicyFile := range bsConfiguration.PolicyFiles {
		policyFile := new(policyFile)
		policyFile.id = bsConfPolicyFile.ID
		policyFile.revisionID = bsConfPolicyFile.RevisionID

		policyFile.name = strings.Join([]string{policyFile.id, "-", policyFile.revisionID}, "")
		policyFile.fileName = strings.Join([]string{policyFile.name, ".tgz"}, "")
		policyFile.tarballURL = bsConfPolicyFile.DownloadURL

		if policyFile.tarballURL != "" {
			url, err := url.Parse(policyFile.tarballURL)
			if err != nil {
				log.Errorf("cannot parse the tarball policy file url: %s [%s]", policyFile.tarballURL, err)
			} else {
				policyFile.queryURL = strings.Join([]string{url.Path[1:], url.RawQuery}, "?")
			}
		}

		policyFile.tarballPath = strings.Join([]string{bsProcess.directoryPath, policyFile.fileName}, "")
		policyFile.folderPath = strings.Join([]string{bsProcess.directoryPath, policyFile.name}, "")

		bsProcess.policyFiles = append(bsProcess.policyFiles, policyFile)
	}
	log.Debug(bsProcess)
}

// downloadPolicyFiles For every policy file, ensure its tarball (downloadable through their download_url) has been downloaded to the server ...
func downloadPolicyFiles(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) {
	log.Debug("downloadPolicyFiles")

	for _, bsPolicyFile := range bsProcess.policyFiles {
		log.Debug("downloading: ", bsPolicyFile.tarballURL)
		_, status, err := bootstrappingSvc.DownloadPolicyFile(bsPolicyFile.queryURL, bsPolicyFile.tarballPath)
		if err != nil {
			log.Errorf("cannot download the tarball policy file: %s [%s]", bsPolicyFile.tarballURL, err)
		}
		if status == 200 {
			bsPolicyFile.downloaded = true
			log.Debug("decompressing: ", bsPolicyFile.tarballPath)
			if err = utils.Untar(bsPolicyFile.tarballPath, bsPolicyFile.folderPath); err != nil {
				log.Errorf("cannot decompress the tarball policy file: %s [%s]", bsPolicyFile.tarballPath, err)
			}
			bsPolicyFile.uncompressed = true
		} else {
			log.Errorf("cannot download the policy file: %v", bsPolicyFile.fileName)
		}
	}
}

// cleanObsoletePolicyFiles cleans off any tarball that is no longer needed.
func cleanObsoletePolicyFiles(bsProcess *bootstrappingProcess) {
	log.Debug("cleanObsoletePolicyFiles")

	// evaluates working folder
	deletableFiles, err := ioutil.ReadDir(bsProcess.directoryPath)
	if err != nil {
		log.Errorf("cannot read directory: %s [%s]", bsProcess.directoryPath, err)
	}

	// builds an array of currently processable files at this looping time
	currentlyProcessableFiles := []string{bsProcess.attributes.fileName} // saved attributes file name
	for _, bsPolicyFile := range bsProcess.policyFiles {
		currentlyProcessableFiles = append(currentlyProcessableFiles, bsPolicyFile.fileName) // Downloaded tgz file names
		currentlyProcessableFiles = append(currentlyProcessableFiles, bsPolicyFile.name)     // Uncompressed folder names
	}

	// removes from deletableFiles array the policy files currently applied
	for _, f := range deletableFiles {
		if !utils.Contains(currentlyProcessableFiles, f.Name()) {
			log.Debug("removing: ", f.Name())
			if err := os.RemoveAll(strings.Join([]string{bsProcess.directoryPath, string(os.PathSeparator), f.Name()}, "")); err != nil {
				log.Errorf("cannot remove: %s [%s]", f.Name(), err)
			}
		}
	}
}

// saveAttributes stores the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
func saveAttributes(bsProcess *bootstrappingProcess) {
	log.Debug("saveAttributes")

	attrs, err := json.Marshal(bsProcess.attributes.rawData)
	if err != nil {
		log.Errorf("cannot process policies attributes: %s [%s]", bsProcess.attributes.revisionID, err)
	}
	if err := ioutil.WriteFile(bsProcess.attributes.filePath, attrs, 0600); err != nil {
		log.Errorf("cannot save policies attributes: %s [%s]", bsProcess.attributes.revisionID, err)
	}
}

// processPolicyFiles applies for each policy the required chef commands, reporting in bunches of N lines
func processPolicyFiles(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) {
	log.Debug("processPolicyFiles")

	for _, bsPolicyFile := range bsProcess.policyFiles {
		command := strings.Join([]string{"cd", bsPolicyFile.folderPath}, " ")
		if runtime.GOOS == "windows" {
			command = strings.Join([]string{command, "SET \"PATH=%PATH%;C:\\ruby\\bin;C:\\opscode\\chef\\bin;C:\\opscode\\chef\\embedded\\bin\""}, ";")
		}
		command = strings.Join([]string{command, strings.Join([]string{"chef-client -z -j", bsProcess.attributes.filePath}, " ")}, ";")
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
		if err != nil {
			log.Errorf("cannot process continuous report command [%s]", err)
		}

		log.Info("completed: ", exitCode)

		bsPolicyFile.executed = exitCode == 0 // policy successfully applied
		//If the command returns with a non-zero value, stop applying policyfiles and continue with the next step.
		if !bsPolicyFile.executed {
			break
		}
	}
}

// reportAppliedConfiguration Inform the platform of applied changes
func reportAppliedConfiguration(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) {
	log.Debug("reportAppliedConfiguration")

	var policyFileRevisionIDs string
	for _, bsPolicyFile := range bsProcess.policyFiles {
		if bsPolicyFile.executed { // only for policies successfully applied
			appliedPolicyMap := map[string]string{bsPolicyFile.id: bsPolicyFile.revisionID}
			appliedPolicyBytes, err := json.Marshal(appliedPolicyMap)
			if err != nil {
				log.Errorf("corrupted candidates policies map [%s]", err)
			}
			policyFileRevisionIDs = strings.Join([]string{policyFileRevisionIDs, string(appliedPolicyBytes)}, "")
		}
	}

	payload := map[string]interface{}{
		"started_at":              bsProcess.startedAt,
		"finished_at":             bsProcess.finishedAt,
		"policyfile_revision_ids": policyFileRevisionIDs,
		"attribute_revision_id":   bsProcess.attributes.revisionID,
	}
	err := bootstrappingSvc.ReportBootstrappingAppliedConfiguration(&payload)
	if err != nil {
		log.Errorf("cannot report applied configuration [%s]", err)
	}
}

// completeBootstrappingSequence evaluates if the first iteration of policies was completed; If case, execute the "scripts boot" command.
func completeBootstrappingSequence(bsProcess *bootstrappingProcess) {
	log.Debug("completeBootstrappingSequence")

	if !allPolicyFilesSuccessfullyApplied {
		checked := true
		for _, bsPolicyFile := range bsProcess.policyFiles {
			if !bsPolicyFile.executed {
				checked = false
				break
			}
		}
		allPolicyFilesSuccessfullyApplied = checked

		if allPolicyFilesSuccessfullyApplied {
			log.Debug("run the boot scripts")
			//run the boot scripts for the server by executing the scripts boot sub-command (as an external process).
			if output, exit, _, _ := utils.RunCmd( strings.Join([]string{os.Args[0], "scripts", "boot"}, " ")); exit != 0 {
				log.Errorf("Error executing scripts boot: (%d) %s", exit, output)
			}
		}
	}
}
