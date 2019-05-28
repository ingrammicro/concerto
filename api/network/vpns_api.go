package network

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// VPNService manages VPN operations
type VPNService struct {
	concertoService utils.ConcertoService
}

// NewVPNService returns a Concerto VPN service
func NewVPNService(concertoService utils.ConcertoService) (*VPNService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &VPNService{
		concertoService: concertoService,
	}, nil
}

// GetVPN returns a VPN by VPC ID
func (dm *VPNService) GetVPN(vpcID string) (vpn *types.Vpn, err error) {
	log.Debug("GetVPN")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/network/vpcs/%s/vpn", vpcID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpn); err != nil {
		return nil, err
	}

	return vpn, nil
}

// CreateVPN creates a VPN
func (dm *VPNService) CreateVPN(vpnVector *map[string]interface{}, vpcID string) (vpn *types.Vpn, err error) {
	log.Debug("CreateVPN")

	data, status, err := dm.concertoService.Post(fmt.Sprintf("/network/vpcs/%s/vpn", vpcID), vpnVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpn); err != nil {
		return nil, err
	}

	return vpn, nil
}

// DeleteVPN deletes VPN by VPC ID
func (dm *VPNService) DeleteVPN(vpcID string) (err error) {
	log.Debug("DeleteVPN")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/network/vpcs/%s/vpn", vpcID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// GetVPNListPlans returns the list of VPN plans for a given VPC ID
func (dm *VPNService) GetVPNListPlans(vpcID string) (vpnPlans []*types.VpnPlan, err error) {
	log.Debug("GetVPNListPlans")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/network/vpcs/%s/vpn_plans", vpcID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpnPlans); err != nil {
		return nil, err
	}

	return vpnPlans, nil
}
