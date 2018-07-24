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

// Time layout
const (
	TimeStampLayout          = "2006-01-02T15:04:05.000000-07:00"
	TimeLayoutYYYYMMDDHHMMSS = "20060102150405"
)

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

// ExecCode saves command into a file and run it
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

// RunFile executes command string contained in a file
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
	output = b.String()

	log.Debugf("Starting Time: %s", startedAt.Format(TimeStampLayout))
	log.Debugf("End Time: %s", finishedAt.Format(TimeStampLayout))
	log.Debugf("Output")
	log.Debugf("")
	log.Debugf("%s", output)
	log.Debugf("")
	log.Infof("Exit Code: %d", exitCode)
	return
}

// RunCmd executes received command string
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

// Save script/command in a temp file
func createCommandWithFilename(command string) (cmd *exec.Cmd, cmdFileName string) {

	cmdFileName = strings.Join([]string{time.Now().Format(TimeLayoutYYYYMMDDHHMMSS), "_", RandomString(10)}, "")
	if runtime.GOOS == "windows" {
		cmdFileName = strings.Join([]string{cmdFileName, ".bat"}, "")
	}

	// Writes content to file
	if err := ioutil.WriteFile(cmdFileName, []byte(command), 0600); err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}

	// Creates command
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdFileName)
	} else {
		cmd = exec.Command("/bin/sh", cmdFileName)
	}
	return
}

// Remove temp file
func deleteTmpCommandFilename(cmdFileName string) {
	err := os.Remove(cmdFileName)
	if err != nil {
		log.Warn("Temp file cannot be removed", err.Error())
	}
}

// RunTracedCmd executes the received command and manages two output pipes (output and error)
// It shouldn't throw any exception/error or stop the process.
func RunTracedCmd(command string) (exitCode int, stdOut string, stdErr string, startedAt time.Time, finishedAt time.Time) {
	log.Debug("RunTracedCmd")

	// Saves script/command in a temp file
	var cmd, cmdFileName = createCommandWithFilename(command)

	// Removes temp file
	defer deleteTmpCommandFilename(cmdFileName)

	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("cmd.StdoutPipe() failed: ", err)
	}

	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		log.Error("cmd.StderrPipe() failed: ", err)
	}

	var errStdout, errStderr error
	var stdoutBuf, stderrBuf bytes.Buffer
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	if err = cmd.Start(); err != nil {
		log.Error("cmd.Start() failed: ", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	if err = cmd.Wait(); err != nil {
		log.Error("cmd.Wait() failed: ", err)
	}

	if errStdout != nil {
		log.Error("failed to capture stdout: ", errStdout)
	}

	if errStderr != nil {
		log.Error("failed to capture stderr: ", errStderr)
	}

	exitCode = extractExitCode(err)
	stdOut = string(stdoutBuf.Bytes())
	stdErr = string(stderrBuf.Bytes())
	startedAt = time.Now()
	finishedAt = time.Now()

	log.Infof("Exit Code: %d", exitCode)
	log.Debugf("Stdout: %s", stdOut)
	log.Debugf("Stderr: %s", stdErr)
	log.Debugf("Starting Time: %s", startedAt.Format(TimeStampLayout))
	log.Debugf("End Time: %s", finishedAt.Format(TimeStampLayout))
	return
}

//RunContinuousCmd executes the received command and process the retrieved response chunks asynchronously.
func RunContinuousCmd(fn func(chunk string) error, command string, thresholdTime int) (int, error) {
	log.Debug("RunContinuousCmd")

	// Saves script/command in a temp file
	var cmd, cmdFileName = createCommandWithFilename(command)

	// Removes temp file
	defer deleteTmpCommandFilename(cmdFileName)

	// Gets the pipe command
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 1, fmt.Errorf("cannot get pipe command %v", err)
	}
	log.Info("==> Executing: ", strings.Join(cmd.Args, " "))

	// Start command asynchronously
	if err = cmd.Start(); err != nil {
		return 1, fmt.Errorf("cannot start the specified command %v", err)
	}

	chunk := ""
	nTime := 0
	timeStart := time.Now()

	scanner := bufio.NewScanner(bufio.NewReader(stdout))
	for scanner.Scan() {
		chunk = strings.Join([]string{chunk, scanner.Text(), "\n"}, "")
		nTime = int(time.Now().Sub(timeStart).Seconds())
		if nTime >= thresholdTime {
			if err := fn(chunk); err != nil {
				nTime = 0
			} else {
				chunk = ""
				nTime = 0
			}
			timeStart = time.Now()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("==> Error: ", err.Error())
		chunk = strings.Join([]string{chunk, err.Error()}, "")
	}

	if len(chunk) > 0 {
		log.Debug("Processing the last pending chunk")
		if err := fn(chunk); err != nil {
			log.Error("Cannot process the last chunk", err.Error())
		}
	}

	err = cmd.Wait()
	exitCode := extractExitCode(err)

	return exitCode, nil
}
