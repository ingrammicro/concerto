package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

// SetProcessIDToFile obtains the process id and save it inside a file
func SetProcessIDToFile(pidFileName string) error {
	log.Debug("SetProcessIDToFile")

	pidValue := os.Getpid()
	log.Debug("current pid:", pidValue)
	if err := ioutil.WriteFile(pidFileName, []byte(strconv.Itoa(pidValue)), 0600); err != nil {
		return err
	}

	return nil
}

// GetProcessIDFromFile reads the process id previously stored in the file
func GetProcessIDFromFile(pidFileName string) (int, error) {
	log.Debug("GetProcessIDFromFile")

	var pid int64

	bytes, err := ioutil.ReadFile(pidFileName)
	if err != nil {
		return 0, err
	}

	pid, err = strconv.ParseInt(string(bytes), 10, 32)
	if err != nil {
		return 0, err
	}
	log.Debug("retrieved pid:", pid)

	return int(pid), nil
}

// StopProcessID stops the process by the given id
func StopProcessID(pid int) error {
	log.Debug("StopProcessID")

	if pid <= 0 {
		return errors.New("invalid pid, a positive value is required")
	}

	log.Debug("getting process:", pid)
	process, err := os.FindProcess(pid)
	if err != nil {
		return errors.New("cannot find the process." + err.Error())
	}

	log.Debug("stopping process:", pid)
	if runtime.GOOS == "windows" {
		err = process.Kill()
		if err != nil {
			return err
		}
	} else {
		if err := process.Signal(syscall.SIGTERM); err != nil {
			return err
		}
	}

	// TODO 20171116 linux arises an error 'waitid: no child processes'
	//log.Debug("waiting for process:", pid)
	//va, err := process.Wait()
	//if err != nil {
	//	return err
	//}
	//log.Debug("stopped:", va.Exited())

	return nil
}

// StopProcess reads the process id from given file and stops the process
func StopProcess(pidFileName string) error {
	log.Debug("StopProcess")

	pid, err := GetProcessIDFromFile(pidFileName)
	if err != nil {
		return errors.New("cannot read the pid file." + err.Error())
	}
	if err := StopProcessID(pid); err != nil {
		return err
	}

	return nil
}
