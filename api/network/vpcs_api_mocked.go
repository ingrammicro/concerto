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

// GetVPCListMocked test mocked function
func GetVPCListMocked(t *testing.T, vpcsIn []*types.Vpc) []*types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcsIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Get", "/network/vpcs").Return(dIn, 200, nil)
	vpcsOut, err := ds.GetVPCList()
	assert.Nil(err, "Error getting VPC list")
	assert.Equal(vpcsIn, vpcsOut, "GetVPCList returned different VPCs")

	return vpcsOut
}

// GetVPCListFailErrMocked test mocked function
func GetVPCListFailErrMocked(t *testing.T, vpcsIn []*types.Vpc) []*types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcsIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Get", "/network/vpcs").Return(dIn, 200, fmt.Errorf("mocked error"))
	vpcsOut, err := ds.GetVPCList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpcsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpcsOut
}

// GetVPCListFailStatusMocked test mocked function
func GetVPCListFailStatusMocked(t *testing.T, vpcsIn []*types.Vpc) []*types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcsIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Get", "/network/vpcs").Return(dIn, 499, nil)
	vpcsOut, err := ds.GetVPCList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpcsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpcsOut
}

// GetVPCListFailJSONMocked test mocked function
func GetVPCListFailJSONMocked(t *testing.T, vpcsIn []*types.Vpc) []*types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/network/vpcs").Return(dIn, 200, nil)
	vpcsOut, err := ds.GetVPCList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpcsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpcsOut
}

// GetVPCMocked test mocked function
func GetVPCMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 200, nil)
	vpcOut, err := ds.GetVPC(vpcIn.ID)
	assert.Nil(err, "Error getting VPC")
	assert.Equal(*vpcIn, *vpcOut, "GetVPC returned different VPCs")

	return vpcOut
}

// GetVPCFailErrMocked test mocked function
func GetVPCFailErrMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	vpcOut, err := ds.GetVPC(vpcIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpcOut
}

// GetVPCFailStatusMocked test mocked function
func GetVPCFailStatusMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 499, nil)
	vpcOut, err := ds.GetVPC(vpcIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpcOut
}

// GetVPCFailJSONMocked test mocked function
func GetVPCFailJSONMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 200, nil)
	vpcOut, err := ds.GetVPC(vpcIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpcOut
}

// CreateVPCMocked test mocked function
func CreateVPCMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Post", "/network/vpcs/", mapIn).Return(dOut, 200, nil)
	vpcOut, err := ds.CreateVPC(mapIn)
	assert.Nil(err, "Error creating VPC list")
	assert.Equal(vpcIn, vpcOut, "CreateVPC returned different VPCs")

	return vpcOut
}

// CreateVPCFailErrMocked test mocked function
func CreateVPCFailErrMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Post", "/network/vpcs/", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	vpcOut, err := ds.CreateVPC(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpcOut
}

// CreateVPCFailStatusMocked test mocked function
func CreateVPCFailStatusMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Post", "/network/vpcs/", mapIn).Return(dOut, 499, nil)
	vpcOut, err := ds.CreateVPC(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpcOut
}

// CreateVPCFailJSONMocked test mocked function
func CreateVPCFailJSONMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/network/vpcs/", mapIn).Return(dIn, 200, nil)
	vpcOut, err := ds.CreateVPC(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpcOut
}

// UpdateVPCMocked test mocked function
func UpdateVPCMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID), mapIn).Return(dOut, 200, nil)
	vpcOut, err := ds.UpdateVPC(mapIn, vpcIn.ID)
	assert.Nil(err, "Error updating VPC list")
	assert.Equal(vpcIn, vpcOut, "UpdateVPC returned different VPCs")

	return vpcOut
}

// UpdateVPCFailErrMocked test mocked function
func UpdateVPCFailErrMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	vpcOut, err := ds.UpdateVPC(mapIn, vpcIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpcOut
}

// UpdateVPCFailStatusMocked test mocked function
func UpdateVPCFailStatusMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID), mapIn).Return(dOut, 499, nil)
	vpcOut, err := ds.UpdateVPC(mapIn, vpcIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return vpcOut
}

// UpdateVPCFailJSONMocked test mocked function
func UpdateVPCFailJSONMocked(t *testing.T, vpcIn *types.Vpc) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID), mapIn).Return(dIn, 200, nil)
	vpcOut, err := ds.UpdateVPC(mapIn, vpcIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpcOut
}

// DeleteVPCMocked test mocked function
func DeleteVPCMocked(t *testing.T, vpcIn *types.Vpc) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteVPC(vpcIn.ID)
	assert.Nil(err, "Error deleting VPC")
}

// DeleteVPCFailErrMocked test mocked function
func DeleteVPCFailErrMocked(t *testing.T, vpcIn *types.Vpc) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteVPC(vpcIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteVPCFailStatusMocked test mocked function
func DeleteVPCFailStatusMocked(t *testing.T, vpcIn *types.Vpc) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVPCService(cs)
	assert.Nil(err, "Couldn't load VPC service")
	assert.NotNil(ds, "VPC service not instanced")

	// to json
	dIn, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/vpcs/%s", vpcIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteVPC(vpcIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
