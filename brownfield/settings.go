package brownfield

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

type Settings struct {
	FQDN          string   `json:"fqdn"`
	Hostname      string   `json:"-"`
	SSHPublicKeys []string `json:"ssh_public_keys"`
}

func applyConcertoSettings(cs *utils.HTTPConcertoservice, f format.Formatter) {
	settings, err := obtainSettings(cs)
	if err != nil {
		f.PrintFatal("Cannot obtain settings", err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("TODO: prepare command for windows execution")
	} else {
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
		err = nixScriptTemplate.Execute(tmpfile, settings)
		if err != nil {
			f.PrintFatal("Cannot not instantiate setup script", err)
		}
		err = tmpfile.Close()
		tmpfileName = tmpfile.Name()
		if err != nil {
			f.PrintFatal("Cannot instantiate setup script", err)
		}
		output, err := exec.Command("cat", tmpfileName).Output()
		if err != nil {
			f.PrintFatal("Error happened running setup script", err)
		}
		fmt.Printf("Setup script ran successfully:\n%s", output)
	}
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
	if settings.FQDN == "" {
		err = fmt.Errorf("retrieved empty FQDN for the server")
		return
	}
	settings.Hostname = strings.Split(settings.FQDN, ".")[0]
	return
}

var nixScriptTemplate = template.Must(template.New("configFile").Parse(`#! /bin/bash

## FQDN settings ##
echo "Setting FQDN"
cat /etc/hosts | grep -v '127\.0\.1\.1' | grep -v {{.Hostname}} | grep -v {{.FQDN}} > /tmp/hosts
echo "127.0.1.1    {{.FQDN}}    {{.Hostname}}" >>/tmp/hosts
mv /tmp/hosts /etc/hosts
hostname {{.Hostname}}
hostname >/etc/hostname

if [ -f /etc/sysconfig/network ]; then
  cat /etc/sysconfig/network | grep -v "HOSTNAME=" >/tmp/sysconfig_network
  echo "HOSTNAME={{.FQDN}}" >>/tmp/sysconfig_network
  mv -f /tmp/sysconfig_network /etc/sysconfig/network
fi

## SSH settings ##
mkdir -p $HOME/.ssh
{{range .SSHPublicKeys}}
echo {{.}} >> $HOME/.ssh/authorized_keys
{{end}}

if grep -v "^PubkeyAuthentication[ \t]*yes" /etc/ssh/sshd_config; then
  sed -i -e "s/^#PubkeyAuthentication[ \t]*yes/PubkeyAuthentication yes/g" -e "s/^PubkeyAuthentication[ \t]*no/PubkeyAuthentication yes/g" /etc/ssh/sshd_config
  sed -i 's/root:x:0:0:root:\\/root:\\/sbin\\/nologin/root:x:0:0:root:\\/root:\\/bin\\/bash/' /etc/passwd
  sed -i -e 's/^AllowUsers /#AllowUsers /' -e 's/^PermitRootLogin /#PermitRootLogin /' /etc/ssh/sshd_config
  /etc/init.d/ssh* restart
fi
`))
