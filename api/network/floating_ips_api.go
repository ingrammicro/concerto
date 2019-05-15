package network

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// FloatingIPService manages FloatingIP operations
type FloatingIPService struct {
	concertoService utils.ConcertoService
}

// NewFloatingIPService returns a Concerto FloatingIP service
func NewFloatingIPService(concertoService utils.ConcertoService) (*FloatingIPService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &FloatingIPService{
		concertoService: concertoService,
	}, nil
}

// GetFloatingIPList returns the list of FloatingIPs as an array of FloatingIP
func (dm *FloatingIPService) GetFloatingIPList(serverID string) (floatingIPs []*types.FloatingIP, err error) {
	log.Debug("GetFloatingIPList")

	path := "/network/floating_ips"
	if serverID != "" {
		path = strings.Join([]string{path, "?server_id=", serverID}, "")
	}
	data, status, err := dm.concertoService.Get(path)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIPs); err != nil {
		return nil, err
	}

	return floatingIPs, nil
}

// GetFloatingIP returns a FloatingIP by its ID
func (dm *FloatingIPService) GetFloatingIP(ID string) (floatingIP *types.FloatingIP, err error) {
	log.Debug("GetFloatingIP")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/network/floating_ips/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// CreateFloatingIP creates a FloatingIP
func (dm *FloatingIPService) CreateFloatingIP(floatingIPVector *map[string]interface{}) (floatingIP *types.FloatingIP, err error) {
	log.Debug("CreateFloatingIP")

	data, status, err := dm.concertoService.Post("/network/floating_ips/", floatingIPVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// UpdateFloatingIP updates a FloatingIP by its ID
func (dm *FloatingIPService) UpdateFloatingIP(floatingIPVector *map[string]interface{}, ID string) (floatingIP *types.FloatingIP, err error) {
	log.Debug("UpdateFloatingIP")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/network/floating_ips/%s", ID), floatingIPVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// AttachFloatingIP attaches a FloatingIP by its ID
func (dm *FloatingIPService) AttachFloatingIP(floatingIPVector *map[string]interface{}, ID string) (server *types.Server, err error) {
	log.Debug("AttachFloatingIP")

	data, status, err := dm.concertoService.Post(fmt.Sprintf("/network/floating_ips/%s/attached_server", ID), floatingIPVector)
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

// DetachFloatingIP detaches a FloatingIP by its ID
func (dm *FloatingIPService) DetachFloatingIP(ID string) (err error) {
	log.Debug("DetachFloatingIP")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/floating_ips/%s/attached_server", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DeleteFloatingIP deletes a FloatingIP by its ID
func (dm *FloatingIPService) DeleteFloatingIP(ID string) (err error) {
	log.Debug("DeleteFloatingIP")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/floating_ips/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DiscardFloatingIP discards a FloatingIP by its ID
func (dm *FloatingIPService) DiscardFloatingIP(ID string) (err error) {
	log.Debug("DiscardFloatingIP")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/floating_ips/%s/discard", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
