package cmdpolling

import (
	"os"
	"os/signal"
	"sync"
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
	PollingPingTimingInterval = 2
	ProcessIdFile             = "imco-polling.pid"
)

var (
	wg               = &sync.WaitGroup{}
	commandProcessed = make(chan bool, 1)
)

// Handle signals
func handleSysSignals() {
	log.Debug("handleSysSignals")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Debug("Ending, signal detected:", <-gracefulStop)
	wg.Done()
}

// Start the polling process
func cmdStart(c *cli.Context) error {
	log.Debug("cmdStart")

	formatter := format.GetFormatter()
	if err := utils.SetProcessIdToFile(ProcessIdFile); err != nil {
		formatter.PrintFatal("cannot create the pid file", err)
	}

	go handleSysSignals()

	wg.Add(1)
	go pingRoutine(c)
	wg.Wait()

	return nil
}

// Stop the polling process
func cmdStop(c *cli.Context) error {
	log.Debug("cmdStop")

	formatter := format.GetFormatter()
	if err := utils.StopProcess(ProcessIdFile); err != nil {
		formatter.PrintFatal("cannot stop the polling process", err)
	}

	log.Info("concerto polling successfully stopped")
	return nil
}

// Main polling background routine
func pingRoutine(c *cli.Context) {
	log.Debug("pingRoutine")

	defer wg.Done()

	formatter := format.GetFormatter()
	pollingSvc := cmd.WireUpPolling(c)

	isRunningCommandRoutine := false
	t := time.NewTicker(PollingPingTimingInterval * time.Second)
	for {
		log.Debug("Requesting for candidate commands status")
		ping, status, err := pollingSvc.Ping()
		if err != nil {
			formatter.PrintError("Couldn't receive polling ping data", err)
		}

		// One command is available, and no process running
		if status == 201 && ping.PendingCommands && !isRunningCommandRoutine {
			log.Debug("Detected a candidate command")
			isRunningCommandRoutine = true
			wg.Add(1)
			go processingCommandRoutine(pollingSvc, formatter)
		}

		select {
		case <-commandProcessed:
			isRunningCommandRoutine = false
		default:
		}
		<-t.C
	}
}

// Subsidiary routine for commands processing
func processingCommandRoutine(pollingSvc *polling.PollingService, formatter format.Formatter) {
	log.Debug("processingCommandRoutine")

	defer wg.Done()

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

		// 3. If command successfully executed, then status is propagated to IMCO
		if command.ExitCode == 0 {
			log.Debug("Reporting command execution status")
			commandIn, err := utils.ItemConvertParamsWithTagAsID(*command)
			if err != nil {
				formatter.PrintError("Couldn't send polling command report data; error parsing payload", err)
			}

			_, status, err := pollingSvc.UpdateCommand(commandIn, command.Id)
			if err != nil {
				formatter.PrintError("Couldn't send polling command report data", err)
			}

			if status == 200 {
				log.Debug("Command execution results successfully reported")
			} else {
				log.Error("Cannot report the command execution results")
			}
		} else {
			log.Error("Cannot run the retrieved command")
		}
	} else {
		log.Error("Cannot retrieve the next command")
	}

	commandProcessed <- true
}
