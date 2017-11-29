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

// SetProcessIdToFile obtains the process id and save it inside a file
func SetProcessIdToFile(pidFileName string) error {
	log.Debug("SetProcessIdToFile")

	pidValue := os.Getpid()
	log.Debug("current pid:", pidValue)
	if err := ioutil.WriteFile(pidFileName, []byte(strconv.Itoa(pidValue)), 0600); err != nil {
		return err
	}

	return nil
}

// GetProcessIdFromFile reads the process id previously stored in the file
func GetProcessIdFromFile(pidFileName string) (int, error) {
	log.Debug("GetProcessIdFromFile")

	var pid int64

	if bytes, err := ioutil.ReadFile(pidFileName); err != nil {
		return 0, err
	} else {
		pid, err = strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return 0, err
		}
		log.Debug("retrieved pid:", pid)
	}

	return int(pid), nil
}

// StopProcessId stops the process by the given id
func StopProcessId(pid int) error {
	log.Debug("StopProcessId")

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

	if pid, err := GetProcessIdFromFile(pidFileName); err != nil {
		return errors.New("cannot read the pid file." + err.Error())
	} else {
		if err := StopProcessId(pid); err != nil {
			return err
		}
	}

	return nil
}
