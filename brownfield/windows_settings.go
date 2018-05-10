// +build windows

package brownfield

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

type Settings struct {
	SSHPublicKeys []string `json:"ssh_public_keys"`
}

func applyConcertoSettings(cs *utils.HTTPConcertoservice, f format.Formatter, username, password string) {
	_, err := obtainSettings(cs)
	if err != nil {
		f.PrintFatal("Cannot obtain settings", err)
	}
	err = sendUsernamePassword(cs, username, password)
	if err != nil {
		f.PrintFatal("Cannot send server credentials", err)
	}
	dir, err := ioutil.TempDir("", "brownfield-configure")
	if err != nil {
		f.PrintFatal("Cannot create temp dir", err)
	}
	defer os.RemoveAll(dir) // clean up

	scriptFilePath := fmt.Sprintf("%s\\configure.bat", dir)

	// Writes content to file
	if err := ioutil.WriteFile(scriptFilePath, []byte(scriptTemplate), 0600); err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}

	output, err := exec.Command("cmd", "/C", scriptFilePath).CombinedOutput()
	if err != nil {
		f.PrintFatal("Error happened running setup script", fmt.Errorf("%s: %v", output, err))
	}
	fmt.Printf("Setup script ran successfully\n")
}

func obtainSettings(cs *utils.HTTPConcertoservice) (*Settings, error) {
	// We do not need settings data, but make the API call to log progress on API service log
	_, _, err := cs.Get("/brownfield/settings")
	if err != nil {
		return nil, err
	}
	return &Settings{}, nil
}

func sendUsernamePassword(cs *utils.HTTPConcertoservice, username, password string) error {
	payload := &map[string]interface{}{
		"settings": map[string]interface{}{
			"username":    username,
			"user_passwd": password,
		},
	}
	body, status, err := cs.Put("/brownfield/settings", payload)
	if err != nil {
		return err
	}
	if status == 403 {
		return fmt.Errorf("server responded with 403 code: authentication was not successful")
	}
	if status >= 300 {
		return fmt.Errorf("server responded with %d code: %s", status, string(body))
	}
	return nil
}

var scriptTemplate = strings.Join([]string{
	`winrm quickconfig -q`,
	`winrm set winrm/config/winrs @{MaxMemoryPerShellMB="1024"}`,
	`winrm set winrm/config @{MaxTimeoutms="1800000"}`,
	`winrm set winrm/config/service @{AllowUnencrypted="true"}`,
	`winrm set winrm/config/service/auth @{Basic="true"}`,
	`netsh advfirewall firewall set rule name="Windows Remote Management (HTTP-In)" profile=public protocol=tcp localport=5985 remoteip=localsubnet new remoteip=any`,
}, " && ")
