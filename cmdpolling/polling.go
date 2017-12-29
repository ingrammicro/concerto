package cmdpolling

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"strings"

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
	ProcessIdFile                         = "imco-polling.pid"
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

	// initialization
	isRunningCommandRoutine := false
	longTicker := time.NewTicker(time.Duration(longTimePeriod) * time.Second)
	shortTicker := time.NewTicker(time.Duration(shortTimePeriod) * time.Second)
	currentTicker := longTicker
	for {
		// no need to request until current command ends
		if !isRunningCommandRoutine {
			log.Debug("Requesting for candidate commands status")
			ping, status, err := pollingSvc.Ping()
			if err != nil {
				formatter.PrintError("Couldn't receive polling ping data", err)
				// in low level error, should the ticker set as the the longest time interval?
			} else {
				// One command is available
				if status == 201 {
					if ping.PendingCommands {
						log.Debug("Detected a candidate command")
						isRunningCommandRoutine = true
						// set short interval timing!?
						// - by default this implies next interval timing sil be short.
						// - If command routine has a over delay, long interval timing will be assigned later
						if currentTicker != shortTicker {
							log.Debug("Ticker assigned: Short")
							shortTicker = time.NewTicker(time.Duration(shortTimePeriod) * time.Second)
							currentTicker = shortTicker
						}
						go processingCommandRoutine(pollingSvc, formatter)
					} else {
						// Since no pending command, long interval timing
						log.Debug("Ticker assigned: Long")
						currentTicker = time.NewTicker(time.Duration(longTimePeriod) * time.Second)
					}
				}
			}
		}

		log.Debug("Waiting...", currentTicker)

		select {
		case <-currentTicker.C:
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing polling")
			return
		}

		select {
		case <-commandProcessed:
			isRunningCommandRoutine = false
		default:
			// if command routine is currently running and short interval timing runs out, a long interval timing is assigned
			if isRunningCommandRoutine && currentTicker == shortTicker {
				log.Debug("Ticker re-assigned: Long")
				currentTicker = time.NewTicker(time.Duration(longTimePeriod) * time.Second)
			}
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
		command.ExitCode, command.Stdout, command.Stderr, _, _ = utils.RunTracedCmd(command.Script)

		// 3. then status is propagated to IMCO
		log.Debug("Reporting command execution status")

		commandIn := map[string]interface{}{
			"id":        command.Id,
			"script":    command.Script,
			"stdout":    command.Stdout,
			"stderr":    command.Stderr,
			"exit_code": command.ExitCode,
		}

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
