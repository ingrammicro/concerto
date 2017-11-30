package cmdpolling

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	var cmdArg string
	if c.Args().Present() {
		cmdArg = c.Args().First()
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

	if err := processCommand(pollingSvc, cmdArg, threshold); err != nil {
		formatter.PrintFatal("cannot process continuous report command", err)
	}

	log.Info("completed")
	return nil
}

func processCommand(pollingSvc *polling.PollingService, cmdArg string, threshold Threshold) error {
	log.Debug("processCommand")

	// Saves script/command in a temp file
	cmdFileName := strings.Join([]string{time.Now().Format("20060102150405"), "_", utils.RandomString(10)}, "")
	if runtime.GOOS == "windows" {
		cmdFileName = strings.Join([]string{cmdFileName, ".bat"}, "")
	}

	// Writes content
	if err := ioutil.WriteFile(cmdFileName, []byte(cmdArg), 0600); err != nil {
		log.Fatalf("Error creating temp file : ", err)
	}

	// Creates command
	var newCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		newCmd = exec.Command("cmd", "/C", cmdFileName)
	} else {
		newCmd = exec.Command("/bin/sh", cmdFileName)
	}

	// Removes temp file
	defer func() {
		err := os.Remove(cmdFileName)
		if err != nil {
			log.Warn("Temp file cannot be removed", err.Error())
		}
	}()

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

	chunk := ""
	nLines, nTime, nBytes := 0, 0, 0
	timeStart := time.Now()

	scanner := bufio.NewScanner(bufio.NewReader(stdout))
	for scanner.Scan() {
		chunk = strings.Join([]string{chunk, scanner.Text(), "\n"}, "")
		nLines++
		nTime = int(time.Now().Sub(timeStart).Seconds())
		nBytes = len(chunk)
		if nLines >= threshold.lines || nTime >= threshold.time || nBytes >= threshold.bytes {
			if err := sendChunks(pollingSvc, chunk); err != nil {
				nLines, nTime = 0, 0
				timeStart = time.Now()
			} else {
				chunk = ""
				nLines, nTime, nBytes = 0, 0, 0
				timeStart = time.Now()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("==> Error: ", err.Error())
		chunk = strings.Join([]string{chunk, err.Error()}, "")
	}

	if len(chunk) > 0 {
		log.Debug("Sending the last pending chunk")
		if err := sendChunks(pollingSvc, chunk); err != nil {
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
