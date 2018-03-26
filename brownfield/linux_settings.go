// +build linux darwin

package brownfield

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

type Settings struct {
	SSHPublicKeys []string `json:"ssh_public_keys"`
}

func applyConcertoSettings(cs *utils.HTTPConcertoservice, f format.Formatter, _, _ string) {
	settings, err := obtainSettings(cs)
	if err != nil {
		f.PrintFatal("Cannot obtain settings", err)
	}

	var tmpfileName string
	tmpfile, err := ioutil.TempFile("", "concerto-setup")
	if err != nil {
		f.PrintFatal("Cannot not open tempfile to write setup script", err)
	}
	defer func() {
		if tmpfileName == "" {
			tmpfile.Close()
		}
		os.Remove(tmpfileName)
	}()
	err = scriptTemplate.Execute(tmpfile, settings)
	if err != nil {
		f.PrintFatal("Cannot not instantiate setup script", err)
	}
	err = tmpfile.Close()
	tmpfileName = tmpfile.Name()
	if err != nil {
		f.PrintFatal("Cannot instantiate setup script", err)
	}
	_, err = exec.Command("bash", tmpfileName).Output()
	if err != nil {
		f.PrintFatal("Error happened running setup script", err)
	}
	fmt.Printf("Setup script ran successfully\n")
}

func obtainSettings(cs *utils.HTTPConcertoservice) (settings *Settings, err error) {
	body, status, err := cs.Get("/brownfield/settings")
	if err != nil {
		return
	}
	if status == 403 {
		err = fmt.Errorf("server responded with 403 code: authentication was not successful")
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %s", status, string(body))
		return
	}
	settings = &Settings{}
	err = json.Unmarshal(body, settings)
	if err != nil {
		err = fmt.Errorf("cannot parse as JSON server response %v: %v", string(body), err)
		return
	}
	return
}

var scriptTemplate = template.Must(template.New("configFile").Parse(`#! /bin/bash

## SSH settings ##
mkdir -p $HOME/.ssh
{{range .SSHPublicKeys}}
echo {{.}} >> $HOME/.ssh/authorized_keys
{{end}}


sed -i -e "s/^#PubkeyAuthentication[ \t]*yes/PubkeyAuthentication yes/g" -e "s/^PubkeyAuthentication[ \t]*no/PubkeyAuthentication yes/g" /etc/ssh/sshd_config
sed -i 's/root:x:0:0:root:\\/root:\\/sbin\\/nologin/root:x:0:0:root:\\/root:\\/bin\\/bash/' /etc/passwd
sed -i -e 's/^AllowUsers /#AllowUsers /' -e 's/^PermitRootLogin /#PermitRootLogin /' /etc/ssh/sshd_config
/etc/init.d/ssh* restart
`))
