package cmdpolling

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"encoding/json"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/polling"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

const (
	DefaultPollingPingTimingInterval = 30
	ProcessIdFile                    = "imco-polling.pid"
)

var (
	commandProcessed = make(chan bool, 1)
)

// Handle signals
func handleSysSignals(cancelFunc context.CancelFunc) {
	log.Debug("handleSysSignals")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Debug("Ending, signal detected:", <-gracefulStop)
	cancelFunc()
}

// Returns the full path to the tmp folder joined with pid management file name
func getProcessIdFilePath() string{
	return strings.Join([]string{os.TempDir(),  string(os.PathSeparator), ProcessIdFile}, "")
}

// Start the polling process
func cmdStart(c *cli.Context) error {
	log.Debug("cmdStart")

	formatter := format.GetFormatter()
	if err := utils.SetProcessIdToFile(getProcessIdFilePath()); err != nil {
		formatter.PrintFatal("cannot create the pid file", err)
	}

	pollingPingTimingInterval := c.Int64("time")
	if !(pollingPingTimingInterval > 0) {
		pollingPingTimingInterval = DefaultPollingPingTimingInterval
	}
	log.Debug("Ping time interval:", pollingPingTimingInterval)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSysSignals(cancel)

	pingRoutine(ctx, c, pollingPingTimingInterval)

	return nil
}

// Stop the polling process
func cmdStop(c *cli.Context) error {
	log.Debug("cmdStop")

	formatter := format.GetFormatter()
	if err := utils.StopProcess(getProcessIdFilePath()); err != nil {
		formatter.PrintFatal("cannot stop the polling process", err)
	}

	log.Info("concerto polling successfully stopped")
	return nil
}

// Main polling background routine
func pingRoutine(ctx context.Context, c *cli.Context, pollingPingTimingInterval int64) {
	log.Debug("pingRoutine")

	formatter := format.GetFormatter()
	pollingSvc := cmd.WireUpPolling(c)

	isRunningCommandRoutine := false
	t := time.NewTicker(time.Duration(pollingPingTimingInterval) * time.Second)
	for {
		log.Debug("Requesting for candidate commands status")
		ping, status, err := pollingSvc.Ping()
		if err != nil {
			formatter.PrintError("Couldn't receive polling ping data", err)
		} else {
			// One command is available, and no process running
			if status == 201 && ping.PendingCommands && !isRunningCommandRoutine {
				log.Debug("Detected a candidate command")
				isRunningCommandRoutine = true
				go processingCommandRoutine(pollingSvc, formatter)
			}
		}

		select {
		case <-commandProcessed:
			isRunningCommandRoutine = false
		default:
		}

		select {
		case <-t.C:
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing polling")
			return
		}
	}
}

// Subsidiary routine for commands processing
func processingCommandRoutine(pollingSvc *polling.PollingService, formatter format.Formatter) {
	log.Debug("processingCommandRoutine")

	// 1. Request for the new command available
	log.Debug("Retrieving available command")
	command, status, err := pollingSvc.GetNextCommand()
	if err != nil {
		formatter.PrintError("Couldn't receive polling command candidate data", err)
	}

	// 2. Execute the retrieved command
	if status == 200 {
		log.Debug("Running the retrieved command")
		command.Stdout, command.ExitCode, _, _ = utils.RunCmd(command.Script)
		command.Stderr = ""
		if command.ExitCode != 0 {
			command.Stderr = command.Stdout
			command.Stdout = ""
		}

		// 3. then status is propagated to IMCO
		log.Debug("Reporting command execution status")

		var commandIn map[string]interface{}
		inRec, _ := json.Marshal(command)
		json.Unmarshal(inRec, &commandIn)

		_, status, err := pollingSvc.UpdateCommand(&commandIn, command.Id)
		if err != nil {
			formatter.PrintError("Couldn't send polling command report data", err)
		}

		if status == 200 {
			log.Debug("Command execution results successfully reported")
		} else {
			log.Error("Cannot report the command execution results")
		}
	} else {
		log.Error("Cannot retrieve the next command")
	}

	commandProcessed <- true
}
