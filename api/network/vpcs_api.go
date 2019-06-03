package network

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// VPCService manages VPC operations
type VPCService struct {
	concertoService utils.ConcertoService
}

// NewVPCService returns a Concerto VPC service
func NewVPCService(concertoService utils.ConcertoService) (*VPCService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &VPCService{
		concertoService: concertoService,
	}, nil
}

// GetVPCList returns the list of VPCs as an array of VPC
func (dm *VPCService) GetVPCList() (vpcs []*types.Vpc, err error) {
	log.Debug("GetVPCList")

	data, status, err := dm.concertoService.Get("/network/vpcs")

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpcs); err != nil {
		return nil, err
	}

	return vpcs, nil
}

// GetVPC returns a VPC by its ID
func (dm *VPCService) GetVPC(ID string) (vpc *types.Vpc, err error) {
	log.Debug("GetVPC")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/network/vpcs/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpc); err != nil {
		return nil, err
	}

	return vpc, nil
}

// CreateVPC creates a VPC
func (dm *VPCService) CreateVPC(vpcVector *map[string]interface{}) (vpc *types.Vpc, err error) {
	log.Debug("CreateVPC")

	data, status, err := dm.concertoService.Post("/network/vpcs/", vpcVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpc); err != nil {
		return nil, err
	}

	return vpc, nil
}

// UpdateVPC updates a VPC by its ID
func (dm *VPCService) UpdateVPC(vpcVector *map[string]interface{}, ID string) (vpc *types.Vpc, err error) {
	log.Debug("UpdateVPC")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/network/vpcs/%s", ID), vpcVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpc); err != nil {
		return nil, err
	}

	return vpc, nil
}

// DeleteVPC deletes a VPC by its ID
func (dm *VPCService) DeleteVPC(ID string) (err error) {
	log.Debug("DeleteVPC")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/vpcs/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DiscardVPC discards a VPC by its ID
func (dm *VPCService) DiscardVPC(ID string) (err error) {
	log.Debug("DiscardVPC")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/vpcs/%s/discard", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
