package cloud

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// ServerService manages server operations
type ServerService struct {
	concertoService utils.ConcertoService
}

// NewServerService returns a Concerto server service
func NewServerService(concertoService utils.ConcertoService) (*ServerService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ServerService{
		concertoService: concertoService,
	}, nil
}

// GetServerList returns the list of servers as an array of Server
func (dm *ServerService) GetServerList() (servers []*types.Server, err error) {
	log.Debug("GetServerList")

	data, status, err := dm.concertoService.Get("/cloud/servers")
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

// GetServer returns a server by its ID
func (dm *ServerService) GetServer(ID string) (server *types.Server, err error) {
	log.Debug("GetServer")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/cloud/servers/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// CreateServer creates a server
func (dm *ServerService) CreateServer(serverVector *map[string]interface{}) (server *types.Server, err error) {
	log.Debug("CreateServer")

	data, status, err := dm.concertoService.Post("/cloud/servers/", serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// UpdateServer updates a server by its ID
func (dm *ServerService) UpdateServer(serverVector *map[string]interface{}, ID string) (server *types.Server, err error) {
	log.Debug("UpdateServer")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/cloud/servers/%s", ID), serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// BootServer boots a server by its ID
func (dm *ServerService) BootServer(serverVector *map[string]interface{}, ID string) (server *types.Server, err error) {
	log.Debug("BootServer")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/cloud/servers/%s/boot", ID), serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// RebootServer reboots a server by its ID
func (dm *ServerService) RebootServer(serverVector *map[string]interface{}, ID string) (server *types.Server, err error) {
	log.Debug("RebootServer")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/cloud/servers/%s/reboot", ID), serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// ShutdownServer shuts down a server by its ID
func (dm *ServerService) ShutdownServer(serverVector *map[string]interface{}, ID string) (server *types.Server, err error) {
	log.Debug("ShutdownServer")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/cloud/servers/%s/shutdown", ID), serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// OverrideServer overrides a server by its ID
func (dm *ServerService) OverrideServer(serverVector *map[string]interface{}, ID string) (server *types.Server, err error) {
	log.Debug("OverrideServer")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/cloud/servers/%s/override", ID), serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// DeleteServer deletes a server by its ID
func (dm *ServerService) DeleteServer(ID string) (err error) {
	log.Debug("DeleteServer")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/cloud/servers/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

//======= Events ==========

// GetEventsList returns a list of events by server ID
func (dm *ServerService) GetEventsList(serverID string) (events []*types.Event, err error) {
	log.Debug("GetEventsList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/cloud/servers/%s/events", serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}

//======= Operational Scripts ==========

// GetOperationalScriptsList returns a list of scripts by server ID
func (dm *ServerService) GetOperationalScriptsList(serverID string) (scripts []*types.ScriptChar, err error) {
	log.Debug("GetOperationalScriptsList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/cloud/servers/%s/operational_scripts", serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

// ExecuteOperationalScript executes an operational script by its server ID and the script id
func (dm *ServerService) ExecuteOperationalScript(serverVector *map[string]interface{}, serverID string, scriptID string) (script *types.Event, err error) {
	log.Debug("ExecuteOperationalScript")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/cloud/servers/%s/operational_scripts/%s/execute", serverID, scriptID), serverVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}
