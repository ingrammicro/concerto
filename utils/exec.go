package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	TimeStampLayout          = "2006-01-02T15:04:05.000000-07:00"
	TimeLayoutYYYYMMDDHHMMSS = "20060102150405"
	DefaultThresholdLines    = 10
	DefaultThresholdTime     = 10
	DefaultThresholdBytes    = 500
)

type Thresholds struct {
	Lines int
	Time  int
	Bytes int
}

func extractExitCode(err error) int {
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return err.(*exec.ExitError).Sys().(syscall.WaitStatus).ExitStatus()
		case *os.PathError:
			return 127
		}
	}
	return 0
}

func ExecCode(code string, path string, filename string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {
	var err error
	var tmp *os.File

	if runtime.GOOS == "windows" {
		tmp, err = os.Create(fmt.Sprintf("%s/%s.bat", path, filename))
	} else {
		tmp, err = os.Create(fmt.Sprintf("%s/%s", path, filename))
	}

	if err != nil {
		log.Fatalf("Error creating temp file : ", err)
	}

	defer tmp.Close()

	_, err = tmp.WriteString(code)
	if err != nil {
		log.Fatalf("Error writing to file : ", err)
	}

	os.Chmod(tmp.Name(), 0777)
	if err != nil {
		log.Fatalf("Error changing permision to file : ", err)
	}

	return RunFile(tmp.Name())
}

func RunFile(command string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {

	var cmd *exec.Cmd

	var b bytes.Buffer
	buffer := bufio.NewWriter(&b)

	if runtime.GOOS == "windows" {
		log.Infof("Command: %s", command)
		cmd = exec.Command("cmd", "/C", command)
	} else {
		log.Infof("Command: %s %s", "/bin/sh", command)
		cmd = exec.Command("/bin/sh", command)
	}

	stdout, err := cmd.StdoutPipe()
	CheckError(err)

	stderr, err := cmd.StderrPipe()
	CheckError(err)

	multi := io.MultiReader(stdout, stderr)

	startedAt = time.Now()
	err = cmd.Start()
	CheckError(err)

	io.Copy(buffer, multi)

	//go io.Copy(buffer, stderr)
	//go io.Copy(buffer, stdout)

	err = cmd.Wait()
	finishedAt = time.Now()
	exitCode = extractExitCode(err)

	err = buffer.Flush()
	CheckError(err)

	log.Debugf("Starting Time: %s", startedAt.Format(TimeStampLayout))
	log.Debugf("End Time: %s", finishedAt.Format(TimeStampLayout))
	log.Debugf("Output")
	log.Debugf("")
	log.Debugf("%s", b.String())
	log.Debugf("")
	log.Infof("Exit Code: %d", exitCode)
	return
}

func RunCmd(command string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		log.Infof("Command: %s", command)
		cmd = exec.Command("cmd", "/C", command)
	} else {
		log.Infof("Command: %s %s", "/bin/sh -c", command)
		cmd = exec.Command("/bin/sh", "-c", command)
	}

	startedAt = time.Now()
	bytes, err := cmd.CombinedOutput()
	finishedAt = time.Now()
	output = strings.TrimSpace(string(bytes))
	exitCode = extractExitCode(err)

	log.Debugf("Starting Time: %s", startedAt.Format(TimeStampLayout))
	log.Debugf("End Time: %s", finishedAt.Format(TimeStampLayout))
	log.Debugf("Output")
	log.Debugf("")
	log.Debugf("%s", output)
	log.Debugf("")
	log.Infof("Exit Code: %d", exitCode)
	return
}

func RunContinuousReportRun(fn func(chunk string) error, cmdArg string, thresholds Thresholds) (int, error) {
	log.Debug("RunContinuousReportRun")

	// command thresholds flags
	if !(thresholds.Lines > 0) {
		thresholds.Lines = DefaultThresholdLines
	}
	if !(thresholds.Time > 0) {
		thresholds.Time = DefaultThresholdTime
	}
	if !(thresholds.Bytes > 0) {
		thresholds.Bytes = DefaultThresholdBytes
	}
	log.Debug("Threshold", thresholds)

	// Saves script/command in a temp file
	cmdFileName := strings.Join([]string{time.Now().Format(TimeLayoutYYYYMMDDHHMMSS), "_", RandomString(10)}, "")
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
		return 1, fmt.Errorf("cannot get pipe command %v", err)
	}
	log.Info("==> Executing: ", strings.Join(newCmd.Args, " "))

	// Start command asynchronously
	if err = newCmd.Start(); err != nil {
		return 1, fmt.Errorf("cannot start the specified command %v", err)
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
		if nLines >= thresholds.Lines || nTime >= thresholds.Time || nBytes >= thresholds.Bytes {
			if err := fn(chunk); err != nil {
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
		if err := fn(chunk); err != nil {
			log.Error("Cannot send the last chunk", err.Error())
		}
	}

	err = newCmd.Wait()
	exitCode := extractExitCode(err)

	return exitCode, nil
}
