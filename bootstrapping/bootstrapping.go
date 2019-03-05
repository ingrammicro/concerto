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
	DefaultTimingInterval     = 600 // 600 seconds = 10 minutes
	DefaultRandomMaxThreshold = 6   // minutes
	DefaultThresholdLines     = 10
	ProcessIDFile             = "imco-bootstrapping.pid"
	RetriesNumber             = 5
)

type bootstrappingProcess struct {
	startedAt      string
	finishedAt     string
	policyFiles    []policyFile
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
	logged       bool
}

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

	timingInterval := c.Int64("time")
	if !(timingInterval > 0) {
		timingInterval = DefaultTimingInterval
	}
	// Adds a random value to the given timing interval!
	// Sleep for a configured amount of time plus a random amount of time (10 minutes plus 0 to 5 minutes, for instance)
	timingInterval = timingInterval + int64(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(DefaultRandomMaxThreshold)*60)
	log.Debug("time interval: ", timingInterval)

	thresholdLines := c.Int("lines")
	if !(thresholdLines > 0) {
		thresholdLines = DefaultThresholdLines
	}
	log.Debug("routine lines threshold: ", thresholdLines)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSysSignals(cancel)

	bootstrappingRoutine(ctx, c, timingInterval, thresholdLines)

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
func bootstrappingRoutine(ctx context.Context, c *cli.Context, timingInterval int64, thresholdLines int) {
	log.Debug("bootstrappingRoutine")

	bootstrappingSvc, formatter := cmd.WireUpBootstrapping(c)
	commandProcessed := make(chan bool, 1)
	isRunningCommandRoutine := false
	currentTicker := time.NewTicker(time.Duration(timingInterval) * time.Second)
	for {
		if !isRunningCommandRoutine {
			isRunningCommandRoutine = true
			go processingCommandRoutine(bootstrappingSvc, formatter, thresholdLines, commandProcessed)
		}

		log.Debug("waiting...", currentTicker)

		select {
		case <-commandProcessed:
			isRunningCommandRoutine = false
			currentTicker.Stop()
			currentTicker = time.NewTicker(time.Duration(timingInterval) * time.Second)
			log.Debug("command processed")
		case <-currentTicker.C:
			log.Debug("ticker")
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing bootstrapping")
			return
		}
	}
}

// Subsidiary routine for commands processing
func processingCommandRoutine(bootstrappingSvc *blueprint.BootstrappingService, formatter format.Formatter, thresholdLines int, commandProcessed chan bool) {
	log.Debug("processingCommandRoutine")

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
		}
	}
	commandProcessed <- true
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

		bsProcess.policyFiles = append(bsProcess.policyFiles, *policyFile)
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

//For every policy file, apply them doing the following:
//	* Extract the tarball to a temporal work directory DIR
//	* Run  `cd DIR; chef-client -z -j path/to/attrs-<attribute_revision_id>.json` while sending the stderr and stdout in bunches of 10 lines to the
// platform via `POST /blueprint/bootstrap_logs` (this resource is a copy of POST /command_polling/bootstrap_logs used in the command_polling command).
// If the command returns with a non-zero value, stop applying policy files and continue with the next step.

// TODO On the first iteration that applies successfully all policy files (runs all `chef-client -z` commands obtaining 0 return codes) only, run the boot scripts for the server by executing the `scripts boot` sub-command (as an external process).
// TODO Just a POC, an starging point. To be completed...
func processPolicyFiles(bootstrappingSvc *blueprint.BootstrappingService, bsProcess *bootstrappingProcess) {
	log.Debug("processPolicyFiles")

	// Run  `cd DIR; chef-client -z -j path/to/attrs-<attribute_revision_id>.json` while sending the stderr and stdout in bunches of
	// 10 lines to the platform via `POST /blueprint/bootstrap_logs` (this resource is a copy of POST /command_polling/bootstrap_logs used in
	// the command_polling command). If the command returns with a non-zero value, stop applying policyfiles and continue with the next step.
	for _, bsPolicyFile := range bsProcess.policyFiles {
		command := strings.Join([]string{"cd", bsPolicyFile.folderPath}, " ")
		if runtime.GOOS == "windows" {
			command = strings.Join([]string{command, "SET \"PATH=%PATH%;C:\\ruby\\bin;C:\\opscode\\chef\\bin;C:\\opscode\\chef\\embedded\\bin\""}, ";")
		}
		command = strings.Join([]string{command, strings.Join([]string{"chef-client -z -j", bsProcess.attributes.filePath}, " ")}, ";")
		log.Debug(command)

		// ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** **
		// TODO ** TO BE REMOVED** !!! for debugging purposes, overriding real command
		command = "ping -c 100 8.8.8.8"
		// ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** ** **

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
	}
}

// reportAppliedConfiguration Inform the platform of applied changes via a `PUT /blueprint/applied_configuration` request
//The `policy file_revision_ids` field should have revision ids set only for those policy files successfully applied on the iteration, that is,
// it should not have any values set for those failing and those skipped because of a previous one failing.
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
