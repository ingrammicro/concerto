package storage

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// VolumeService manages Volume operations
type VolumeService struct {
	concertoService utils.ConcertoService
}

// NewVolumeService returns a Concerto Volume service
func NewVolumeService(concertoService utils.ConcertoService) (*VolumeService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &VolumeService{
		concertoService: concertoService,
	}, nil
}

// GetVolumeList returns the list of Volumes as an array of Volume
func (dm *VolumeService) GetVolumeList(serverID string) (volumes []*types.Volume, err error) {
	log.Debug("GetVolumeList")

	path := "/storage/volumes"
	if serverID != "" {
		path = fmt.Sprintf("/cloud/servers/%s/volumes", serverID)

	}
	data, status, err := dm.concertoService.Get(path)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volumes); err != nil {
		return nil, err
	}

	return volumes, nil
}

// GetVolume returns a Volume by its ID
func (dm *VolumeService) GetVolume(volumeID string) (volume *types.Volume, err error) {
	log.Debug("GetVolume")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/storage/volumes/%s", volumeID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}

// CreateVolume creates a Volume
func (dm *VolumeService) CreateVolume(volumeVector *map[string]interface{}) (volume *types.Volume, err error) {
	log.Debug("CreateVolume")

	data, status, err := dm.concertoService.Post("/storage/volumes/", volumeVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}

// UpdateVolume updates a Volume by its ID
func (dm *VolumeService) UpdateVolume(volumeVector *map[string]interface{}, volumeID string) (volume *types.Volume, err error) {
	log.Debug("UpdateVolume")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/storage/volumes/%s", volumeID), volumeVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}

// AttachVolume attaches a Volume by its ID
func (dm *VolumeService) AttachVolume(volumeVector *map[string]interface{}, volumeID string) (server *types.Server, err error) {
	log.Debug("AttachVolume")

	data, status, err := dm.concertoService.Post(fmt.Sprintf("/storage/volumes/%s/attached_server", volumeID), volumeVector)
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

// DetachVolume detaches a Volume by its ID
func (dm *VolumeService) DetachVolume(volumeID string) (err error) {
	log.Debug("DetachVolume")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/storage/volumes/%s/attached_server", volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DeleteVolume deletes a Volume by its ID
func (dm *VolumeService) DeleteVolume(volumeID string) (err error) {
	log.Debug("DeleteVolume")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/storage/volumes/%s", volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DiscardVolume discards a Volume by its ID
func (dm *VolumeService) DiscardVolume(volumeID string) (err error) {
	log.Debug("DiscardVolume")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/storage/volumes/%s/discard", volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
