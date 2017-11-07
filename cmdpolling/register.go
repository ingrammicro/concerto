package cmdpolling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

var configFileTemplate = template.Must(template.New("configFile").Parse(`<concerto version="1.0" server="{{.APIEndpoint}}" log_file="{{.LogFile}}" log_level="{{.LogLevel}}">
<ssl cert="{{.CertPath}}" key="{{.KeyPath}}" server_ca="{{.CaCertPath}}" />
</concerto>
`))

func cmdRegister(c *cli.Context) error {
	f := format.GetFormatter()
	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't read config", err)
	}
	if !config.CurrentUserIsAdmin {
		if runtime.GOOS == "windows" {
			f.PrintFatal("Must run as administrator user", fmt.Errorf("running as non-administrator user"))
		} else {
			f.PrintFatal("Must run as super-user", fmt.Errorf("running as non-administrator user"))
		}
	}
	rootCACert, cert, key, err := obtainServerKeys(config)
	if err != nil {
		f.PrintFatal("Couldn't obtain server keys", err)
	}
	err = configureServerKeys(config, rootCACert, cert, key)
	if err != nil {
		f.PrintFatal("Couldn't configure server keys", err)
	}
	fmt.Printf("Concerto agent successfully registered, configuration file placed at %s\n", config.ConfFile)
	return nil
}

func obtainServerKeys(config *utils.Config) (rootCAcert string, cert string, key string, err error) {
	cs, err := utils.NewHTTPConcertoServiceWithCommandPolling(config)
	if err != nil {
		return
	}
	payload := make(map[string]interface{})
	body, status, err := cs.Post("/command_polling/api_key", &payload)
	if err != nil {
		return
	}
	if status == 403 {
		err = fmt.Errorf("server responded with 403 code: the polling token is not valid, maybe it expired")
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %s", status, string(body))
		return
	}
	responseData := make(map[string]interface{})
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	irootCAcert, ok := responseData["root_ca_cert"]
	if !ok {
		err = fmt.Errorf("server response did not include root CA cert: %v", responseData)
		return
	}
	rootCAcert, ok = irootCAcert.(string)
	if !ok {
		err = fmt.Errorf("server response returned a %T as root CA cert, expected a string", irootCAcert)
		return
	}
	iCert, ok := responseData["cert"]
	if !ok {
		err = fmt.Errorf("server response did not include server cert: %v", responseData)
		return
	}
	cert, ok = iCert.(string)
	if !ok {
		err = fmt.Errorf("server response returned a %T as server cert, expected a string", iCert)
		return
	}
	iKey, ok := responseData["key"]
	if !ok {
		err = fmt.Errorf("server response did not include server private key: %v", responseData)
		return
	}
	key, ok = iKey.(string)
	if !ok {
		err = fmt.Errorf("server response returned a %T as server private key, expected a string", iKey)
	}
	return
}

func configureServerKeys(config *utils.Config, rootCACert, cert, key string) error {
	configFileData := &struct {
		APIEndpoint string
		LogFile     string
		LogLevel    string
		CertPath    string
		KeyPath     string
		CaCertPath  string
	}{config.APIEndpoint, config.LogFile, config.LogLevel,
		config.Certificate.Cert, config.Certificate.Key, config.Certificate.Ca}

	if configFileData.LogLevel == "" {
		configFileData.LogLevel = "info"
	}
	if configFileData.LogFile == "" {
		configFileData.LogFile = utils.GetDefaultLogFilePath()
	}
	if configFileData.CaCertPath == "" {
		configFileData.CaCertPath = utils.GetDefaultCaCertFilePath()
	}
	if configFileData.CertPath == "" {
		configFileData.CertPath = utils.GetDefaultCertFilePath()
	}
	if configFileData.KeyPath == "" {
		configFileData.KeyPath = utils.GetDefaultKeyFilePath()
	}
	err := os.MkdirAll(filepath.Dir(configFileData.CaCertPath), 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place root CA cert: %v", err)
	}
	err = ioutil.WriteFile(configFileData.CaCertPath, []byte(rootCACert), 0644)
	if err != nil {
		return fmt.Errorf("cannot write root CA cert: %v", err)
	}
	err = os.MkdirAll(filepath.Dir(configFileData.CertPath), 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place server cert: %v", err)
	}
	err = ioutil.WriteFile(configFileData.CertPath, []byte(cert), 0644)
	if err != nil {
		return fmt.Errorf("cannot write server cert: %v", err)
	}
	err = os.MkdirAll(filepath.Dir(configFileData.KeyPath), 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place server key: %v", err)
	}
	err = ioutil.WriteFile(configFileData.KeyPath, []byte(key), 0600)
	if err != nil {
		return fmt.Errorf("cannot write server key: %v", err)
	}
	err = os.MkdirAll(config.ConfLocation, 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place config file: %v", err)
	}
	f, err := os.OpenFile(config.ConfFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Could not open config file for writing: %v", err)
	}
	defer f.Close()
	err = configFileTemplate.Execute(f, configFileData)
	if err != nil {
		return fmt.Errorf("Could not generate config file contents: %v", err)
	}
	return nil
}
