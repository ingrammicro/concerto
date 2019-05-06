package utils

import (
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/go-homedir"
)

const windowsServerConfigFile = "c:\\cio\\client.xml"
const nixServerConfigFile = "/etc/cio/client.xml"
const defaultConcertoEndpoint = "https://clients.concerto.io:886/"

const windowsServerLogFilePath = "c:\\cio\\log\\concerto-client.log"
const windowsServerCaCertPath = "c:\\cio\\client_ssl\\ca_cert.pem"
const windowsServerCertPath = "c:\\cio\\client_ssl\\cert.pem"
const windowsServerKeyPath = "c:\\cio\\client_ssl\\private\\key.pem"
const nixServerLogFilePath = "/var/log/concerto-client.log"
const nixServerCaCertPath = "/etc/cio/client_ssl/ca_cert.pem"
const nixServerCertPath = "/etc/cio/client_ssl/cert.pem"
const nixServerKeyPath = "/etc/cio/client_ssl/private/key.pem"

// Config stores configuration file contents
type Config struct {
	XMLName             xml.Name `xml:"concerto"`
	APIEndpoint         string   `xml:"server,attr"`
	LogFile             string   `xml:"log_file,attr"`
	LogLevel            string   `xml:"log_level,attr"`
	Certificate         Cert     `xml:"ssl"`
	ConfLocation        string
	ConfFile            string
	IsHost              bool
	ConcertoURL         string
	BrownfieldToken     string
	CommandPollingToken string
	ServerID            string
	CurrentUserName     string
	CurrentUserIsAdmin  bool
}

// Cert stores cert files location
type Cert struct {
	Cert string `xml:"cert,attr"`
	Key  string `xml:"key,attr"`
	Ca   string `xml:"server_ca,attr"`
}

var cachedConfig *Config

// GetConcertoConfig returns concerto configuration
func GetConcertoConfig() (*Config, error) {
	if cachedConfig == nil {
		return nil, fmt.Errorf("configuration hasn't been initialized")
	}
	return cachedConfig, nil
}

// InitializeConcertoConfig creates the concerto configuration structure
func InitializeConcertoConfig(c *cli.Context) (*Config, error) {
	log.Debug("InitializeConcertoConfig")
	if cachedConfig != nil {
		return cachedConfig, nil
	}

	cachedConfig = &Config{}

	if err := cachedConfig.readBrownfieldToken(c); err != nil {
		return nil, err
	}

	if err := cachedConfig.readCommandPollingConfig(c); err != nil {
		return nil, err
	}

	// where config file must me
	if err := cachedConfig.evaluateConcertoConfigFile(c); err != nil {
		return nil, err
	}

	// read config contents
	log.Debugf("Reading configuration from %s", cachedConfig.ConfFile)
	if err := cachedConfig.readConcertoConfig(c); err != nil {
		return nil, err
	}

	// add login URL. Needed for setup
	if err := cachedConfig.readConcertoURL(); err != nil {
		return nil, err
	}

	// check if isHost. Needed to show appropriate options
	if err := cachedConfig.evaluateCertificate(); err != nil {
		return nil, err
	}

	// evaluates API endpoint url
	if err := cachedConfig.evaluateAPIEndpointURL(); err != nil {
		return nil, err
	}

	debugShowConfig()
	return cachedConfig, nil
}

func debugShowConfig() {
	if log.GetLevel() < log.DebugLevel {
		return
	}

	if cachedConfig == nil {
		log.Debug("Concerto configuration not loaded")
	}

	debugStruct("", *cachedConfig)
	// c := reflect.ValueOf(*cachedConfig)
	// for i := 0; i < c.NumField(); i++ {
	// 	if c.Type().Field(i).Type.String() != "xml.Name" {
	// 		log.WithField(c.Type().Field(i).Name, c.Field(i).Interface()).Debug("Configuration item")
	// 	}
	// }
}

// debugStruct iterates struct and show in debug console all items and subitems
func debugStruct(prefix string, item interface{}) {
	c := reflect.ValueOf(item)
	for i := 0; i < c.NumField(); i++ {
		if c.Type().Field(i).Type.String() != "xml.Name" {

			name := c.Type().Field(i).Name
			value := c.Field(i).Interface()

			// if value is struct, iterate with recursion
			if c.Type().Field(i).Type.Kind() == reflect.Struct {
				debugStruct(name, value)
			} else {
				if prefix != "" {
					name = fmt.Sprintf("%s.%s", prefix, name)
				}
				log.WithField(name, value).Debug("Configuration item")
			}
		}
	}
}

// IsAgentMode returns whether CLI is acting as Server Or Client mode
func (config *Config) IsAgentMode() bool {
	return config.IsHost || config.BrownfieldToken != "" || config.CommandPollingToken != ""
}

// IsConfigReady returns whether configurations items are filled
func (config *Config) IsConfigReady() bool {
	if config.APIEndpoint == "" ||
		config.Certificate.Cert == "" ||
		config.Certificate.Key == "" ||
		config.Certificate.Ca == "" {
		return false
	}
	return true
}

// IsConfigReadySetup returns whether we can use setup command
func (config *Config) IsConfigReadySetup() bool {
	return config.ConcertoURL != ""
}

// IsConfigReadyBrownfield returns whether config is ready for brownfield token
// authentication
func (config *Config) IsConfigReadyBrownfield() bool {
	if config.APIEndpoint == "" ||
		config.BrownfieldToken == "" {
		return false
	}
	return true
}

// IsConfigReadyCommandPolling returns whether config is ready for polling token
// authentication
func (config *Config) IsConfigReadyCommandPolling() bool {
	if config.APIEndpoint == "" ||
		config.CommandPollingToken == "" ||
		config.ServerID == "" {
		return false
	}
	return true
}

// readConcertoConfig reads Concerto config file located at fileLocation
func (config *Config) readConcertoConfig(c *cli.Context) error {
	log.Debug("Reading Concerto Configuration")
	if FileExists(config.ConfFile) {
		// file exists, read it's contents

		xmlFile, err := os.Open(config.ConfFile)
		if err != nil {
			return err
		}
		defer xmlFile.Close()
		b, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			return fmt.Errorf("configuration File %s couldn't be read", config.ConfFile)
		}

		if err = xml.Unmarshal(b, &config); err != nil {
			return fmt.Errorf("configuration File %s does not have valid XML format", config.ConfFile)
		}

	} else {
		log.Debugf("Configuration File %s does not exist. Reading environment variables", config.ConfFile)
	}

	// overwrite with environment/arguments vars
	if overwEP := c.String("concerto-endpoint"); overwEP != "" {
		log.Debug("Concerto APIEndpoint taken from env/args")
		config.APIEndpoint = overwEP
	}

	if overwCert := c.String("client-cert"); overwCert != "" {
		log.Debug("Certificate path taken from env/args")
		config.Certificate.Cert = overwCert
	}

	if overwKey := c.String("client-key"); overwKey != "" {
		log.Debug("Certificate key path taken from env/args")
		config.Certificate.Key = overwKey
	}

	if overwCa := c.String("ca-cert"); overwCa != "" {
		log.Debug("CA certificate path taken from env/args")
		config.Certificate.Ca = overwCa
	}

	// if endpoint empty set default
	// we can't set the default from flags, because it would overwrite config file
	if config.APIEndpoint == "" {
		config.APIEndpoint = defaultConcertoEndpoint
	}

	return nil
}

func (config *Config) evaluateCurrentUser() (*user.User, error) {
	currUser, err := user.Current()
	if err != nil {
		log.Debugf("Couldn't use os.user to get user details: %s", err.Error())
		dir, err := homedir.Dir()
		if err != nil {
			return nil, fmt.Errorf("couldn't get home dir for current user: %s", err.Error())
		}
		currUser = &user.User{
			Username: getUsername(),
			HomeDir:  dir,
		}
	}
	if runtime.GOOS == "windows" {
		currUser.Username = currUser.Username[strings.LastIndex(currUser.Username, "\\")+1:]
		log.Debugf("Windows username is %s", currUser.Username)
		config.CurrentUserIsAdmin = currUser.Gid == "S-1-5-32-544" || isWinAdministrator(currUser.Username) || canPerformAdministratorTasks()
	} else {
		config.CurrentUserIsAdmin = currUser.Uid == "0" || currUser.Username == "root"
	}
	config.CurrentUserName = currUser.Username
	return currUser, nil
}

// evaluateConcertoConfigFile returns path to concerto config file
func (config *Config) evaluateConcertoConfigFile(c *cli.Context) error {
	log.Debug("evaluateConcertoConfigFile")
	currUser, err := config.evaluateCurrentUser()
	if err != nil {
		return err
	}
	if configFile := c.String("concerto-config"); configFile != "" {

		log.Debug("Concerto configuration file location taken from env/args")
		config.ConfFile = configFile

	} else {

		if runtime.GOOS == "windows" {
			if config.CurrentUserIsAdmin && (config.BrownfieldToken != "" || (config.CommandPollingToken != "" && config.ServerID != "") || FileExists(windowsServerConfigFile)) {
				log.Debugf("Current user is administrator, setting config file as %s", windowsServerConfigFile)
				config.ConfFile = windowsServerConfigFile
			} else {
				// User mode Windows
				log.Debugf("Current user is regular user: %s", currUser.Username)
				config.ConfFile = filepath.Join(currUser.HomeDir, ".concerto/client.xml")
			}
		} else {
			// Server mode *nix
			if config.CurrentUserIsAdmin && (config.BrownfieldToken != "" || (config.CommandPollingToken != "" && config.ServerID != "") || FileExists(nixServerConfigFile)) {
				config.ConfFile = nixServerConfigFile
			} else {
				// User mode *nix
				config.ConfFile = filepath.Join(currUser.HomeDir, ".concerto/client.xml")
			}
		}
	}
	config.ConfLocation = path.Dir(config.ConfFile)
	return nil
}

// getUsername gets username by env variable.
// os.user is dependant on cgo, so cross compiling won't work
func getUsername() string {
	log.Debug("getUsername")
	u := "unknown"
	osUser := ""

	switch runtime.GOOS {
	case "darwin", "linux", "solaris":
		osUser = os.Getenv("USER")
	case "windows":
		osUser = os.Getenv("USERNAME")

		// remove domain
		osUser = osUser[strings.LastIndex(osUser, "\\")+1:]
		log.Debugf("Windows user has been transformed into %s", osUser)

		// HACK ugly ... if localized administrator, translate to administrator
		if isWinAdministrator(osUser) {
			osUser = "Administrator"
		}
	}

	if osUser != "" {
		u = osUser
	}
	return u
}

func isWinAdministrator(user string) bool {
	return user == "Järjestelmänvalvoja" ||
		user == "Administrateur" ||
		user == "Rendszergazda" ||
		user == "Administrador" ||
		user == "Администратор" ||
		user == "Administratör" ||
		user == "Administrator" ||
		user == "SYSTEM" ||
		user == "imco"
}

func canPerformAdministratorTasks() bool {
	f, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

// readConcertoURL reads URL from CONCERTO_URL environment or calculates using API URL
func (config *Config) readConcertoURL() error {

	if config.ConcertoURL != "" {
		return nil
	}

	if overwURL := os.Getenv("CONCERTO_URL"); overwURL != "" {
		config.ConcertoURL = overwURL
		log.Debug("Concerto URL taken from CONCERTO_URL")
		return nil
	}

	cURL, err := url.Parse(config.APIEndpoint)
	if err != nil {
		return err
	}

	tokenHost := strings.Split(cURL.Host, ":")
	tokenFqdn := strings.Split(tokenHost[0], ".")

	if !strings.Contains(cURL.Host, "staging") {
		tokenFqdn[0] = "start"
	}

	config.ConcertoURL = fmt.Sprintf("%s://%s/", cURL.Scheme, strings.Join(tokenFqdn, "."))
	return nil
}

// readConcertoURL reads URL from CONCERTO_URL environment or calculates using API URL
func (config *Config) readBrownfieldToken(c *cli.Context) error {
	if config.BrownfieldToken != "" {
		return nil
	}

	// overwrite with environment/arguments vars
	if overwBrownfieldToken := c.String("concerto-brownfield-token"); overwBrownfieldToken != "" {
		log.Debug("Concerto Brownfield token taken from env/args")
		config.BrownfieldToken = overwBrownfieldToken
	}

	return nil
}

func (config *Config) readCommandPollingConfig(c *cli.Context) error {
	if config.CommandPollingToken != "" || config.ServerID != "" {
		return nil
	}

	// overwrite with environment/arguments vars
	if overwCommandPollingToken := c.String("concerto-command-polling-token"); overwCommandPollingToken != "" {
		log.Debug("Concerto Command Polling token taken from env/args")
		config.CommandPollingToken = overwCommandPollingToken
	}

	if overwServerID := c.String("concerto-server-id"); overwServerID != "" {
		log.Debug("Concerto Server ID taken from env/args")
		config.ServerID = overwServerID
	}

	return nil
}

// evaluateCertificate determines if a certificate has been issued for a host
func (config *Config) evaluateCertificate() error {

	if FileExists(config.Certificate.Cert) {

		data, err := ioutil.ReadFile(config.Certificate.Cert)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(data)

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}

		if len(cert.Subject.OrganizationalUnit) > 0 {
			if cert.Subject.OrganizationalUnit[0] == "Hosts" {
				config.IsHost = true
				return nil
			}
		} else if len(cert.Issuer.Organization) > 0 {
			if cert.Issuer.Organization[0] == "Tapp" {
				config.IsHost = true
				return nil
			}
		}
	}
	config.IsHost = false
	return nil
}

// evaluateAPIEndpointURL evaluates if API endpoint url is valid, advising if invalid version defined, and adapting if required
func (config *Config) evaluateAPIEndpointURL() error {
	log.Debug("evaluateAPIEndpointURL")

	// remove ending slash if exist
	config.APIEndpoint = strings.TrimRight(config.APIEndpoint, "/")

	// In User mode, endpoint url should include API version
	if !config.IsAgentMode() {
		cURL, err := url.Parse(config.APIEndpoint)
		if err != nil {
			return err
		}
		if cURL.Path == "" {
			config.APIEndpoint = strings.Join([]string{config.APIEndpoint, VERSION_API_USER_MODE}, "/")
			log.Warnf("Defined API server endpoint url does not include API version. Normalized to latest version (%s): %s", VERSION_API_USER_MODE, config.APIEndpoint)
		} else if cURL.Path != strings.Join([]string{"/", VERSION_API_USER_MODE}, "") {
			log.Warnf("Defined API server endpoint url does not match the latest supported API version (%s). Found %s", VERSION_API_USER_MODE, cURL.Path)
		}
	}

	return nil
}

// GetDefaultLogFilePath returns default concerto configuration path file
func GetDefaultLogFilePath() string {
	if runtime.GOOS == "windows" {
		return windowsServerLogFilePath
	}
	return nixServerLogFilePath
}

// GetDefaultCaCertFilePath returns default concerto configuration path file
func GetDefaultCaCertFilePath() string {
	if runtime.GOOS == "windows" {
		return windowsServerCaCertPath
	}
	return nixServerCaCertPath
}

// GetDefaultCertFilePath returns default concerto configuration path file
func GetDefaultCertFilePath() string {
	if runtime.GOOS == "windows" {
		return windowsServerCertPath
	}
	return nixServerCertPath
}

// GetDefaultKeyFilePath returns default concerto configuration path file
func GetDefaultKeyFilePath() string {
	if runtime.GOOS == "windows" {
		return windowsServerKeyPath
	}
	return nixServerKeyPath
}
