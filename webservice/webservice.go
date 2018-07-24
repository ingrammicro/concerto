package webservice

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/utils"
)

// WebService stores Web Service data
type WebService struct {
	config *utils.Config
	client *http.Client
}

const contentDispositionRegex = "filename=\\\"([^\\\"]*){1}\\\""

// NewWebService creates a new Web Service
func NewWebService() (*WebService, error) {
	config, err := utils.GetConcertoConfig()
	if err != nil {
		return nil, err
	}

	if !config.IsConfigReady() {
		return nil, fmt.Errorf("Configuration is incomplete")
	}

	client, err := httpClient(config)
	if err != nil {
		return nil, err
	}

	return &WebService{config, client}, nil
}

func httpClient(config *utils.Config) (*http.Client, error) {

	// Loads Clients Certificates and creates and 509KeyPair
	cert, err := tls.LoadX509KeyPair(config.Certificate.Cert, config.Certificate.Key)
	if err != nil {
		return nil, err
	}

	// Creates a client with specific transport configurations
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	return client, nil
}

// Post sends a POST to the specified HTTP endpoint,
func (w *WebService) Post(endpoint string, json []byte) ([]byte, int, error) {
	log.Debugf("Connecting: %s%s", w.config.APIEndpoint, endpoint)
	output := strings.NewReader(string(json))
	response, err := w.client.Post(w.config.APIEndpoint+endpoint, "application/json", output)

	log.Debugf("Posting: %s", output)
	if err != nil {
		return nil, 4000, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Debugf("Response: %s", body)
	log.Debugf("Status code: %s", response.Status)

	return body, response.StatusCode, nil
}

// Put sends a PUT to the specified HTTP endpoint,
func (w *WebService) Put(endpoint string, json []byte) ([]byte, int, error) {
	log.Debugf("Connecting: %s%s", w.config.APIEndpoint, endpoint)
	output := strings.NewReader(string(json))

	request, err := http.NewRequest("PUT", w.config.APIEndpoint+endpoint, output)
	if err != nil {
		return nil, -1, err
	}

	request.Header = map[string][]string{"Content-type": {"application/json"}}
	response, err := w.client.Do(request)

	log.Debugf("Putting: %s", endpoint)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Debugf("Response: %s", body)
	log.Debugf("Status code: %s", response.Status)
	return body, response.StatusCode, nil
}

// Delete sends a DELETE to the specified HTTP endpoint,
func (w *WebService) Delete(endpoint string) ([]byte, int, error) {
	log.Debugf("Connecting: %s%s", w.config.APIEndpoint, endpoint)

	request, err := http.NewRequest("DELETE", w.config.APIEndpoint+endpoint, nil)
	if err != nil {
		return nil, -1, err
	}

	log.Debugf("Deleting: %s", endpoint)
	response, err := w.client.Do(request)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Debugf("Response: %s", body)
	log.Debugf("Status code: %s", response.Status)
	return body, response.StatusCode, nil
}

// Get sends a GET to the specified HTTP endpoint,
func (w *WebService) Get(endpoint string) ([]byte, int, error) {

	log.Debugf("Connecting: %s%s", w.config.APIEndpoint, endpoint)
	response, err := w.client.Get(w.config.APIEndpoint + endpoint)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()

	log.Debugf("Status code: %s", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, -1, err
	}

	log.Debugf("Response: %s", string(body))
	return body, response.StatusCode, nil
}

// GetFile sends a GET to the specified HTTP endpoint, Downloading the retrieved data to a file
func (w *WebService) GetFile(endpoint string, directoryPath string) (string, error) {

	log.Debugf("Connecting: %s%s", w.config.APIEndpoint, endpoint)
	response, err := w.client.Get(w.config.APIEndpoint + endpoint)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	log.Debugf("Status code: %s", response.Status)

	r, err := regexp.Compile(contentDispositionRegex)
	if err != nil {
		return "", err
	}

	fileName := r.FindStringSubmatch(response.Header.Get("Content-Disposition"))[1]
	if err != nil {
		return "", err
	}
	realFileName := fmt.Sprintf("%s/%s", directoryPath, fileName)

	output, err := os.Create(realFileName)
	if err != nil {
		return "", err
	}
	defer output.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return "", err
	}

	log.Debugf("%#v bytes downloaded", n)
	return realFileName, nil
}
