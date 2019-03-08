package cmdpolling

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/polling"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

const (
	DefaultPollingPingTimingIntervalLong  = 30
	DefaultPollingPingTimingIntervalShort = 5
	ProcessIdFile                         = "cio-polling.pid"
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
func getProcessIdFilePath() string {
	return strings.Join([]string{os.TempDir(), string(os.PathSeparator), ProcessIdFile}, "")
}

// Start the polling process
func cmdStart(c *cli.Context) error {
	log.Debug("cmdStart")

	formatter := format.GetFormatter()
	if err := utils.SetProcessIdToFile(getProcessIdFilePath()); err != nil {
		formatter.PrintFatal("cannot create the pid file", err)
	}

	pollingPingTimingIntervalLong := c.Int64("longTime")
	if !(pollingPingTimingIntervalLong > 0) {
		pollingPingTimingIntervalLong = DefaultPollingPingTimingIntervalLong
	}
	log.Debug("Ping long time interval:", pollingPingTimingIntervalLong)

	pollingPingTimingIntervalShort := c.Int64("shortTime")
	if !(pollingPingTimingIntervalShort > 0) {
		pollingPingTimingIntervalShort = DefaultPollingPingTimingIntervalShort
	}
	log.Debug("Ping short time interval:", pollingPingTimingIntervalShort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSysSignals(cancel)

	pingRoutine(ctx, c, pollingPingTimingIntervalLong, pollingPingTimingIntervalShort)

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
func pingRoutine(ctx context.Context, c *cli.Context, longTimePeriod int64, shortTimePeriod int64) {
	log.Debug("pingRoutine")

	formatter := format.GetFormatter()
	pollingSvc := cmd.WireUpPolling(c)
	commandProcessed := make(chan bool, 1)

	// initialization
	isRunningCommandRoutine := false
	longTicker := time.NewTicker(time.Duration(longTimePeriod) * time.Second)
	currentTicker := longTicker
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
				go processingCommandRoutine(pollingSvc, formatter, commandProcessed)
			}
		}

		log.Debug("Waiting...", currentTicker)

		select {
		case <-commandProcessed:
			isRunningCommandRoutine = false
			if currentTicker != longTicker {
				currentTicker.Stop()
			}
			log.Debug("Ticker assigned: short")
			currentTicker = time.NewTicker(time.Duration(shortTimePeriod) * time.Second)
		case <-currentTicker.C:
			if currentTicker != longTicker {
				currentTicker.Stop()
				log.Debug("Ticker assigned: Long")
				currentTicker = longTicker
			}
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing polling")
			return
		}
	}
}

// Subsidiary routine for commands processing
func processingCommandRoutine(pollingSvc *polling.PollingService, formatter format.Formatter, commandProcessed chan bool) {
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
		command.ExitCode, command.Stdout, command.Stderr, _, _ = utils.RunTracedCmd(command.Script)

		// 3. then status is propagated to IMCO
		log.Debug("Reporting command execution status")

		commandIn := map[string]interface{}{
			"id":        command.ID,
			"script":    command.Script,
			"stdout":    command.Stdout,
			"stderr":    command.Stderr,
			"exit_code": command.ExitCode,
		}

		_, status, err := pollingSvc.UpdateCommand(&commandIn, command.ID)
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
