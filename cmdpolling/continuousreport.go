package cmdpolling

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/polling"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

const (
	DefaultThresholdLines = 10
	DefaultThresholdTime  = 10
	DefaultThresholdBytes = 500
	RetriesNumber         = 5
	RetriesFactor         = 3
)

type Threshold struct {
	lines int
	time  int
	bytes int
}

func cmdContinuousReportRun(c *cli.Context) error {
	log.Debug("cmdContinuousReportRun")

	formatter := format.GetFormatter()
	pollingSvc := cmd.WireUpPolling(c)

	// cli command argument
	var commandArg string
	if c.Args().Present() {
		commandArg = c.Args().First()
	} else {
		formatter.PrintFatal("argument missing", errors.New("a script or command is required"))
	}

	// cli command thresholds flags
	thresholdLines := c.Int("lines")
	if !(thresholdLines > 0) {
		thresholdLines = DefaultThresholdLines
	}
	thresholdTime := c.Int("time")
	if !(thresholdTime > 0) {
		thresholdTime = DefaultThresholdTime
	}
	thresholdBytes := c.Int("bytes")
	if !(thresholdBytes > 0) {
		thresholdBytes = DefaultThresholdBytes
	}
	threshold := Threshold{lines: thresholdLines, time: thresholdTime, bytes: thresholdBytes}
	log.Debug("Threshold", threshold)

	if err := processCommand(pollingSvc, commandArg, threshold); err != nil {
		formatter.PrintFatal("cannot process continuous report command", err)
	}

	log.Info("completed")
	return nil
}

func processCommand(pollingSvc *polling.PollingService, commandArg string, threshold Threshold) error {
	log.Debug("processCommand")

	var newCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		newCmd = exec.Command("cmd", "/C", commandArg)
	} else {
		newCmd = exec.Command("/bin/sh", "-c", commandArg)
	}

	// Gets the pipe command
	stdout, err := newCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("cannot get pipe command %v", err)
	}
	log.Info("==> Executing: ", strings.Join(newCmd.Args, " "))

	// Start command asynchronously
	if err = newCmd.Start(); err != nil {
		return fmt.Errorf("cannot start the specified command %v", err)
	}

	line := ""
	nLines, nTime, nBytes := 0, 0, 0
	timeStart := time.Now()

	scanner := bufio.NewScanner(bufio.NewReader(stdout))
	for scanner.Scan() {
		line += scanner.Text()
		if len(line) > 0 {
			nLines++
			nTime = int(time.Now().Sub(timeStart).Seconds())
			nBytes = len(line)
			if nLines >= threshold.lines || nTime >= threshold.time || nBytes >= threshold.bytes {
				if err := sendChunks(pollingSvc, line); err != nil {
					return err
				} else {
					line = ""
					nLines, nTime, nBytes = 0, 0, 0
					timeStart = time.Now()
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("==> Error: ", err.Error())
		if err := sendChunks(pollingSvc, "error reading standard input:"+err.Error()); err != nil {
			return err
		}
	} else {
		log.Debug("Sending the last pending chunk")
		if err := sendChunks(pollingSvc, line); err != nil {
			return err
		}
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

func sendChunks(pollingSvc *polling.PollingService, chunk string) error {
	log.Debug("sendChunks")

	err := retry(RetriesNumber, time.Second, func() error {
		log.Debug("Sending: ", chunk)
		command := types.PollingContinuousReport{
			Stdout: chunk,
		}
		commandIn, err := utils.ItemConvertParamsWithTagAsID(command)
		if err != nil {
			return fmt.Errorf("error parsing payload %v", err)
		}

		_, statusCode, err := pollingSvc.ReportBootstrapLog(commandIn)
		switch {
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
