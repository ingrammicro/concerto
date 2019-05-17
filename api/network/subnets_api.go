package network

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// SubnetService manages Subnet operations
type SubnetService struct {
	concertoService utils.ConcertoService
}

// NewSubnetService returns a Concerto Subnet service
func NewSubnetService(concertoService utils.ConcertoService) (*SubnetService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &SubnetService{
		concertoService: concertoService,
	}, nil
}

// GetSubnetList returns the list of Subnets of a VPC as an array of Subnet
func (dm *SubnetService) GetSubnetList(vpcID string) (subnets []*types.Subnet, err error) {
	log.Debug("GetSubnetList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/network/vpcs/%s/subnets", vpcID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnets); err != nil {
		return nil, err
	}

	return subnets, nil
}

// GetSubnet returns a Subnet by its ID
func (dm *SubnetService) GetSubnet(ID string) (subnet *types.Subnet, err error) {
	log.Debug("GetSubnet")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/network/subnets/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnet); err != nil {
		return nil, err
	}

	return subnet, nil
}

// CreateSubnet creates a Subnet
func (dm *SubnetService) CreateSubnet(subnetVector *map[string]interface{}, vpcID string) (subnet *types.Subnet, err error) {
	log.Debug("CreateSubnet")

	data, status, err := dm.concertoService.Post(fmt.Sprintf("/network/vpcs/%s/subnets", vpcID), subnetVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnet); err != nil {
		return nil, err
	}

	return subnet, nil
}

// UpdateSubnet updates a Subnet by its ID
func (dm *SubnetService) UpdateSubnet(subnetVector *map[string]interface{}, ID string) (subnet *types.Subnet, err error) {
	log.Debug("UpdateSubnet")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/network/subnets/%s", ID), subnetVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnet); err != nil {
		return nil, err
	}

	return subnet, nil
}

// DeleteSubnet deletes a Subnet by its ID
func (dm *SubnetService) DeleteSubnet(ID string) (err error) {
	log.Debug("DeleteSubnet")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/subnets/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
