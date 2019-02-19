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

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
	"fmt"
)

const (
	// DefaultTimingInterval Default period for looping
	DefaultTimingInterval     = 600 // 600 seconds = 10 minutes
	DefaultRandomMaxThreshold = 6   // minutes

	// ProcessIDFile
	ProcessIDFile = "imco-bootstrapping.pid"

	RetriesNumber        = 5
	RetriesFactor        = 3
	DefaultThresholdTime = 10
)

type bootstrappingStatus struct {
	startedAt      string
	finishedAt     string
	policiesStatus []policyStatus
	attributes     attributesStatus
}
type attributesStatus struct {
	revisionID string
	filename   string
	filePath   string
	rawData    *json.RawMessage
}

type policyStatus struct {
	id          string
	revisionID  string
	name        string
	filename    string
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
	log.Debug("time interval:", timingInterval)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSysSignals(cancel)

	bootstrappingRoutine(ctx, c, timingInterval)

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
func bootstrappingRoutine(ctx context.Context, c *cli.Context, timingInterval int64) {
	log.Debug("bootstrappingRoutine")

	//formatter := format.GetFormatter()
	bootstrappingSvc, formatter := cmd.WireUpBootstrapping(c)
	commandProcessed := make(chan bool, 1)

	// initialization
	currentTicker := time.NewTicker(time.Duration(timingInterval) * time.Second)
	for {
		go processingCommandRoutine(bootstrappingSvc, formatter, commandProcessed)

		log.Debug("Waiting...", currentTicker)

		select {
		//case <-commandProcessed:
		//	isRunningCommandRoutine = false
		//	log.Debug("commandProcessed")
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
func processingCommandRoutine(bootstrappingSvc *blueprint.BootstrappingService, formatter format.Formatter, commandProcessed chan bool) {
	log.Debug("processingCommandRoutine")

	// Inquire about desired configuration changes to be applied by querying the `GET /blueprint/configuration` endpoint. This will provide a JSON response with the desired configuration changes
	bsConfiguration, status, err := bootstrappingSvc.GetBootstrappingConfiguration()
	if err != nil {
		formatter.PrintError("Couldn't receive bootstrapping data", err)
	} else {
		if status == 200 {
			bsStatus := new(bootstrappingStatus)
			directoryPath := getProcessingFolderFilePath()

			// proto structures
			if err := initializePrototype(directoryPath, bsConfiguration, bsStatus); err != nil {
				formatter.PrintError("Cannot initialize the policy files prototypes", err)
			}

			// TODO Currently as a step previous to process tarballs policies but this can be done as a part or processing, and using defer for removing files (tgz & folder!?)
			// For every policyFile, ensure its tarball (downloadable through their download_url) has been downloaded to the server ...
			if err := downloadPolicyFiles(bootstrappingSvc, bsStatus); err != nil {
				formatter.PrintError("Cannot download the policy files", err)
			}

			//... and clean off any tarball that is no longer needed.
			if err := cleanObsoletePolicyFiles(directoryPath, bsStatus); err != nil {
				formatter.PrintError("Cannot clean obsolete policy files", err)
			}

			// Store the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
			if err := saveAttributes(bsStatus); err != nil {
				formatter.PrintError("Cannot save policy files attributes ", err)
			}

			// Process tarballs policies
			if err := processPolicyFiles(bootstrappingSvc, bsStatus); err != nil {
				formatter.PrintError("Cannot process policy files ", err)
			}

			// Inform the platform of applied changes via a `PUT /blueprint/applied_configuration` request with a JSON payload similar to
			reportAppliedConfiguration(bootstrappingSvc, bsStatus)
		}
	}

	// TODO
	commandProcessed <- true
}

func initializePrototype(directoryPath string, bsConfiguration *types.BootstrappingConfiguration, bsStatus *bootstrappingStatus) error {
	log.Debug("initializePrototype")

	log.Debug("Initializing bootstrapping structures")

	bsStatus.startedAt = time.Now().UTC().String() // TODO UTC?

	// Attributes
	bsStatus.attributes.revisionID = bsConfiguration.AttributeRevisionID
	bsStatus.attributes.filename = strings.Join([]string{"attrs-", bsStatus.attributes.revisionID, ".json"}, "")
	bsStatus.attributes.filePath = strings.Join([]string{directoryPath, bsStatus.attributes.filename}, "")
	bsStatus.attributes.rawData = bsConfiguration.Attributes

	// Policies
	for _, policyFile := range bsConfiguration.PolicyFiles {
		policyStatus := new(policyStatus)
		policyStatus.id = policyFile.ID
		policyStatus.revisionID = policyFile.RevisionID
		policyStatus.name = strings.Join([]string{policyFile.ID, "-", policyFile.RevisionID}, "")
		policyStatus.filename = strings.Join([]string{policyStatus.name, ".tgz"}, "")
		policyStatus.tarballURL = policyFile.DownloadURL

		url, err := url.Parse(policyStatus.tarballURL)
		if err != nil {
			// TODO should it be an error?
			return err
		}
		policyStatus.queryURL = strings.Join([]string{url.Path[1:], url.RawQuery}, "?")

		policyStatus.tarballPath = strings.Join([]string{directoryPath, policyStatus.filename}, "")
		policyStatus.folderPath = strings.Join([]string{directoryPath, policyStatus.name}, "")

		bsStatus.policiesStatus = append(bsStatus.policiesStatus, *policyStatus)
	}
	log.Debug(bsStatus)
	return nil
}

// downloadPolicyFiles For every policy file, ensure its tarball (downloadable through their download_url) has been downloaded to the server ...
func downloadPolicyFiles(bootstrappingSvc *blueprint.BootstrappingService, bsStatus *bootstrappingStatus) error {
	log.Debug("downloadPolicyFiles")

	for _, policyStatus := range bsStatus.policiesStatus {
		log.Debug("Downloading: ", policyStatus.tarballURL)
		_, status, err := bootstrappingSvc.DownloadPolicyFile(policyStatus.queryURL, policyStatus.tarballPath)
		if err != nil {
			return err
		}
		if status == 200 {
			policyStatus.downloaded = true
			log.Debug("Uncompressing: ", policyStatus.tarballPath)
			if err = utils.Untar(policyStatus.tarballPath, policyStatus.folderPath); err != nil {
				return err
			}
			policyStatus.uncompressed = true
		} else {
			// TODO should it be an error?
			log.Error("Cannot download the policy file: ", policyStatus.filename)
		}
	}
	return nil
}

// cleanObsoletePolicyFiles cleans off any tarball that is no longer needed.
func cleanObsoletePolicyFiles(directoryPath string, bsStatus *bootstrappingStatus) error {
	log.Debug("cleanObsoletePolicyFiles")

	// builds an array of currently processable files at this looping time
	currentlyProcessableFiles := []string{bsStatus.attributes.filename} // saved attributes file name
	for _, policyStatus := range bsStatus.policiesStatus {
		currentlyProcessableFiles = append(currentlyProcessableFiles, policyStatus.filename) // Downloaded tgz file names
		currentlyProcessableFiles = append(currentlyProcessableFiles, policyStatus.name)     // Uncompressed folder names
	}

	// evaluates working folder
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		// TODO should it be an error?
		log.Warn("Cannot read directory: ", directoryPath, err)
	}

	// removes files not regarding to any of current policy files
	for _, f := range files {
		if !utils.Contains(currentlyProcessableFiles, f.Name()) {
			log.Debug("Removing: ", f.Name())
			if err := utils.RemoveFileInfo(f, strings.Join([]string{directoryPath, string(os.PathSeparator), f.Name()}, "")); err != nil {
				// TODO should it be an error?
				log.Warn("Cannot remove: ", f.Name(), err)
			}
		}
	}
	return nil // TODO should it be managed as error?
}

// saveAttributes stores the attributes as JSON in a file with name `attrs-<attribute_revision_id>.json`
func saveAttributes(bsStatus *bootstrappingStatus) error {
	log.Debug("saveAttributes")

	attrs, err := json.Marshal(bsStatus.attributes.rawData)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(bsStatus.attributes.filePath, attrs, 0600); err != nil {
		return err
	}
	return nil
}


//For every policy file, apply them doing the following:
//	* Extract the tarball to a temporal work directory DIR
//	* Run  `cd DIR; chef-client -z -j path/to/attrs-<attribute_revision_id>.json` while sending the stderr and stdout in bunches of 10 lines to the
// platform via `POST /blueprint/bootstrap_logs` (this resource is a copy of POST /command_polling/bootstrap_logs used in the command_polling command).
// If the command returns with a non-zero value, stop applying policy files and continue with the next step.

// TODO On the first iteration that applies successfully all policy files (runs all `chef-client -z` commands obtaining 0 return codes) only, run the boot scripts for the server by executing the `scripts boot` sub-command (as an external process).
// TODO Just a POC, an starging point. To be completed...
func processPolicyFiles(bootstrappingSvc *blueprint.BootstrappingService, bsStatus *bootstrappingStatus) error {
	log.Debug("processPolicyFiles")

	// Run  `cd DIR; chef-client -z -j path/to/attrs-<attribute_revision_id>.json` while sending the stderr and stdout in bunches of
	// 10 lines to the platform via `POST /blueprint/bootstrap_logs` (this resource is a copy of POST /command_polling/bootstrap_logs used in
	// the command_polling command). If the command returns with a non-zero value, stop applying policyfiles and continue with the next step.
	for _, policyStatus := range bsStatus.policiesStatus {
		log.Warn(policyStatus.folderPath)

		// TODO cd <policyStatus.folderPath>; chef-client -z -j <bsStatus.attributes.filePath>`
		command := "ping -c 100 8.8.8.8"

		// cli command threshold flag
		thresholdTime := DefaultThresholdTime
		log.Debug("Time threshold: ", thresholdTime)

		// Custom method for chunks processing
		fn := func(chunk string) error {
			log.Debug("sendChunks")
			err := retry(RetriesNumber, time.Second, func() error {
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

		// TODO This method was implemented in some moment based on nLines, nTime, bBytes? Currently only working with thresholdTime
		exitCode, err := utils.RunContinuousCmd(fn, command, thresholdTime)
		if err != nil {
			log.Error("cannot process continuous report command", err)
		}

		log.Info("completed: ", exitCode)
	}
	return nil
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	log.Debug("retry")

	if err := fn(); err != nil {
		if attempts--; attempts > 0 {
			log.Debug("Waiting to retry: ", sleep)
			time.Sleep(sleep)
			return retry(attempts, RetriesFactor*sleep, fn)
		}
		return err
	}
	return nil
}

// reportAppliedConfiguration Inform the platform of applied changes via a `PUT /blueprint/applied_configuration` request
//The `policy file_revision_ids` field should have revision ids set only for those policy files successfully applied on the iteration, that is,
// it should not have any values set for those failing and those skipped because of a previous one failing.
func reportAppliedConfiguration(bootstrappingSvc *blueprint.BootstrappingService, bsStatus *bootstrappingStatus) error {
	log.Debug("reportAppliedConfiguration")

	bsStatus.finishedAt = time.Now().UTC().String() // TODO UTC?

	var policyfileRevisionIDs string
	for _, policyStatus := range bsStatus.policiesStatus {
		if policyStatus.executed { // only for policies successfully applied
			appliedPolicyMap := map[string]string{policyStatus.id: policyStatus.revisionID}
			appliedPolicyBytes, err := json.Marshal(appliedPolicyMap)
			if err != nil {
				// TODO should it be an error?
				return err
			}
			policyfileRevisionIDs = strings.Join([]string{policyfileRevisionIDs, string(appliedPolicyBytes)}, "")
		}
	}

	payload := map[string]interface{}{
		"started_at":              bsStatus.startedAt,
		"finished_at":             bsStatus.finishedAt,
		"policyfile_revision_ids": policyfileRevisionIDs,
		"attribute_revision_id":   bsStatus.attributes.revisionID,
	}
	err := bootstrappingSvc.ReportBootstrappingAppliedConfiguration(&payload)
	if err != nil {
		// TODO should it be an error?
		return err
	}
	return nil
}
