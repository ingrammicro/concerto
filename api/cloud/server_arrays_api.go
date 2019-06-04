package cloud

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// ServerArrayService manages server array operations
type ServerArrayService struct {
	concertoService utils.ConcertoService
}

// NewServerArrayService returns a Concerto server array service
func NewServerArrayService(concertoService utils.ConcertoService) (*ServerArrayService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ServerArrayService{
		concertoService: concertoService,
	}, nil
}

// GetServerArrayList returns the list of server arrays as an array of ServerArray
func (sas *ServerArrayService) GetServerArrayList() (serverArrays []*types.ServerArray, err error) {
	log.Debug("GetServerArrayList")

	data, status, err := sas.concertoService.Get("/cloud/server_arrays")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArrays); err != nil {
		return nil, err
	}

	return serverArrays, nil
}

// GetServerArray returns a server array by its ID
func (sas *ServerArrayService) GetServerArray(serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("GetServerArray")

	data, status, err := sas.concertoService.Get(fmt.Sprintf("/cloud/server_arrays/%s", serverArrayID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// CreateServerArray creates a server array
func (sas *ServerArrayService) CreateServerArray(serverArrayVector *map[string]interface{}) (serverArray *types.ServerArray, err error) {
	log.Debug("CreateServerArray")

	data, status, err := sas.concertoService.Post("/cloud/server_arrays/", serverArrayVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// UpdateServerArray updates a server array by its ID
func (sas *ServerArrayService) UpdateServerArray(serverArrayVector *map[string]interface{}, serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("UpdateServerArray")

	data, status, err := sas.concertoService.Put(fmt.Sprintf("/cloud/server_arrays/%s", serverArrayID), serverArrayVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// BootServerArray boots a server array by its ID
func (sas *ServerArrayService) BootServerArray(serverArrayVector *map[string]interface{}, serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("BootServerArray")

	data, status, err := sas.concertoService.Put(fmt.Sprintf("/cloud/server_arrays/%s/boot", serverArrayID), serverArrayVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// ShutdownServerArray shuts down a server array by its ID
func (sas *ServerArrayService) ShutdownServerArray(serverArrayVector *map[string]interface{}, serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("ShutdownServerArray")

	data, status, err := sas.concertoService.Put(fmt.Sprintf("/cloud/server_arrays/%s/shutdown", serverArrayID), serverArrayVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// EmptyServerArray empties a server array by its ID
func (sas *ServerArrayService) EmptyServerArray(serverArrayVector *map[string]interface{}, serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("EmptyServerArray")

	data, status, err := sas.concertoService.Put(fmt.Sprintf("/cloud/server_arrays/%s/empty", serverArrayID), serverArrayVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// EnlargeServerArray enlarges a server array by its ID
func (sas *ServerArrayService) EnlargeServerArray(serverArrayVector *map[string]interface{}, serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("EnlargeServerArray")

	data, status, err := sas.concertoService.Post(fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayID), serverArrayVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// GetServerArrayServerList returns the list of servers in a server array as an array of Server
func (sas *ServerArrayService) GetServerArrayServerList(serverArrayID string) (servers []*types.Server, err error) {
	log.Debug("GetServerArrayServerList")

	data, status, err := sas.concertoService.Get(fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &servers); err != nil {
		return nil, err
	}

	return servers, nil
}

// DeleteServerArray deletes a server array by its ID
func (sas *ServerArrayService) DeleteServerArray(serverArrayID string) (err error) {
	log.Debug("DeleteServerArray")

	data, status, err := sas.concertoService.Delete(fmt.Sprintf("/cloud/server_arrays/%s", serverArrayID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
