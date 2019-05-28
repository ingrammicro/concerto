package network

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetVPNMocked test mocked function
func GetVPNMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 200, nil)
	vpnOut, err := ds.GetVPN(vpnIn.ID)
	assert.Nil(err, "Error getting VPN")
	assert.Equal(*vpnIn, *vpnOut, "GetVPN returned different VPNs")

	return vpnOut
}

// GetVPNFailErrMocked test mocked function
func GetVPNFailErrMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	vpnOut, err := ds.GetVPN(vpnIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpnOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpnOut
}

// GetVPNFailStatusMocked test mocked function
func GetVPNFailStatusMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 499, nil)
	vpnOut, err := ds.GetVPN(vpnIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpnOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpnOut
}

// GetVPNFailJSONMocked test mocked function
func GetVPNFailJSONMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 200, nil)
	vpnOut, err := ds.GetVPN(vpnIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpnOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpnOut
}

// CreateVPNMocked test mocked function
func CreateVPNMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// to json
	dOut, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.VpcID), mapIn).Return(dOut, 200, nil)
	vpnOut, err := ds.CreateVPN(mapIn, vpnIn.VpcID)
	assert.Nil(err, "Error creating VPN list")
	assert.Equal(vpnIn, vpnOut, "CreateVPN returned different VPNs")

	return vpnOut
}

// CreateVPNFailErrMocked test mocked function
func CreateVPNFailErrMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// to json
	dOut, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.VpcID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	vpnOut, err := ds.CreateVPN(mapIn, vpnIn.VpcID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpnOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpnOut
}

// CreateVPNFailStatusMocked test mocked function
func CreateVPNFailStatusMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// to json
	dOut, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.VpcID), mapIn).Return(dOut, 499, nil)
	vpnOut, err := ds.CreateVPN(mapIn, vpnIn.VpcID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpnOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpnOut
}

// CreateVPNFailJSONMocked test mocked function
func CreateVPNFailJSONMocked(t *testing.T, vpnIn *types.Vpn) *types.Vpn {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.VpcID), mapIn).Return(dIn, 200, nil)
	vpnOut, err := ds.CreateVPN(mapIn, vpnIn.VpcID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpnOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpnOut
}

// DeleteVPNMocked test mocked function
func DeleteVPNMocked(t *testing.T, vpnIn *types.Vpn) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteVPN(vpnIn.ID)
	assert.Nil(err, "Error deleting VPN")
}

// DeleteVPNFailErrMocked test mocked function
func DeleteVPNFailErrMocked(t *testing.T, vpnIn *types.Vpn) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteVPN(vpnIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteVPNFailStatusMocked test mocked function
func DeleteVPNFailStatusMocked(t *testing.T, vpnIn *types.Vpn) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/vpcs/%s/vpn", vpnIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteVPN(vpnIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// GetVPNListPlansMocked test mocked function
func GetVPNListPlansMocked(t *testing.T, vpnPlansIn []*types.VpnPlan, vpcID string) []*types.VpnPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnPlansIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn_plans", vpcID)).Return(dIn, 200, nil)
	vpnPlansOut, err := ds.GetVPNListPlans(vpcID)
	assert.Nil(err, "Error getting VPN plans list")
	assert.Equal(vpnPlansIn, vpnPlansOut, "GetVPNListPlans returned different VPN plans")

	return vpnPlansOut
}

// GetVPNListPlansFailErrMocked test mocked function
func GetVPNListPlansFailErrMocked(t *testing.T, vpnPlansIn []*types.VpnPlan, vpcID string) []*types.VpnPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnPlansIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn_plans", vpcID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	vpnPlansOut, err := ds.GetVPNListPlans(vpcID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpnPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpnPlansOut
}

// GetVPNListPlansFailStatusMocked test mocked function
func GetVPNListPlansFailStatusMocked(t *testing.T, vpnPlansIn []*types.VpnPlan, vpcID string) []*types.VpnPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// to json
	dIn, err := json.Marshal(vpnPlansIn)
	assert.Nil(err, "VPN test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn_plans", vpcID)).Return(dIn, 499, nil)
	vpnPlansOut, err := ds.GetVPNListPlans(vpcID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpnPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpnPlansOut
}

// GetVPNListPlansFailJSONMocked test mocked function
func GetVPNListPlansFailJSONMocked(t *testing.T, vpnPlansIn []*types.VpnPlan, vpcID string) []*types.VpnPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPNService(cs)
	assert.Nil(err, "Couldn't load VPN service")
	assert.NotNil(ds, "VPN service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/vpn_plans", vpcID)).Return(dIn, 200, nil)
	vpnPlansOut, err := ds.GetVPNListPlans(vpcID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpnPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpnPlansOut
}
