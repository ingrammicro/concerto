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

// GetFloatingIPListMocked test mocked function
func GetFloatingIPListMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", "/network/floating_ips").Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIPList("")
	assert.Nil(err, "Error getting floating IP list")
	assert.Equal(floatingIPIn, floatingIPOut, "GetFloatingIPList returned different floating IPs")

	return floatingIPOut
}

func GetFloatingIPListMockedFilteredByServer(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/floating_ips?server_id=%s", floatingIPIn[0].AttachedServerID)).Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIPList(floatingIPIn[0].AttachedServerID)
	assert.Nil(err, "Error getting floating IP list filtered by server")
	assert.Equal(floatingIPIn, floatingIPOut, "GetFloatingIPList returned different floating IPs")

	return floatingIPOut
}

// GetFloatingIPListFailErrMocked test mocked function
func GetFloatingIPListFailErrMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", "/network/floating_ips").Return(dIn, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.GetFloatingIPList("")

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// GetFloatingIPListFailStatusMocked test mocked function
func GetFloatingIPListFailStatusMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", "/network/floating_ips").Return(dIn, 499, nil)
	floatingIPOut, err := ds.GetFloatingIPList("")

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// GetFloatingIPListFailJSONMocked test mocked function
func GetFloatingIPListFailJSONMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/network/floating_ips").Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIPList("")

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// GetFloatingIPMocked test mocked function
func GetFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error getting floating IP")
	assert.Equal(*floatingIPIn, *floatingIPOut, "GetFloatingIP returned different floating IPs")

	return floatingIPOut
}

// GetFloatingIPFailErrMocked test mocked function
func GetFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// GetFloatingIPFailStatusMocked test mocked function
func GetFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 499, nil)
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// GetFloatingIPFailJSONMocked test mocked function
func GetFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// CreateFloatingIPMocked test mocked function
func CreateFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", "/network/floating_ips/", mapIn).Return(dOut, 200, nil)
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)
	assert.Nil(err, "Error creating floating IP list")
	assert.Equal(floatingIPIn, floatingIPOut, "CreateFloatingIP returned different floating IPs")

	return floatingIPOut
}

// CreateFloatingIPFailErrMocked test mocked function
func CreateFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", "/network/floating_ips/", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// CreateFloatingIPFailStatusMocked test mocked function
func CreateFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", "/network/floating_ips/", mapIn).Return(dOut, 499, nil)
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// CreateFloatingIPFailJSONMocked test mocked function
func CreateFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/network/floating_ips/", mapIn).Return(dIn, 200, nil)
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// UpdateFloatingIPMocked test mocked function
func UpdateFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID), mapIn).Return(dOut, 200, nil)
	floatingIPOut, err := ds.UpdateFloatingIP(mapIn, floatingIPIn.ID)
	assert.Nil(err, "Error updating floating IP list")
	assert.Equal(floatingIPIn, floatingIPOut, "UpdateFloatingIP returned different floating IPs")

	return floatingIPOut
}

// UpdateFloatingIPFailErrMocked test mocked function
func UpdateFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.UpdateFloatingIP(mapIn, floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// UpdateFloatingIPFailStatusMocked test mocked function
func UpdateFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID), mapIn).Return(dOut, 499, nil)
	floatingIPOut, err := ds.UpdateFloatingIP(mapIn, floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return floatingIPOut
}

// UpdateFloatingIPFailJSONMocked test mocked function
func UpdateFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID), mapIn).Return(dIn, 200, nil)
	floatingIPOut, err := ds.UpdateFloatingIP(mapIn, floatingIPIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// AttachFloatingIPMocked test mocked function
func AttachFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: floatingIPIn.AttachedServerID})
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.AttachFloatingIP(mapIn, floatingIPIn.ID)
	assert.Nil(err, "Error attaching floating IP")
	assert.Equal(floatingIPIn.AttachedServerID, serverOut.ID, "AttachFloatingIP returned invalid values")

	return serverOut
}

// AttachFloatingIPFailErrMocked test mocked function
func AttachFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: floatingIPIn.AttachedServerID})
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.AttachFloatingIP(mapIn, floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// AttachFloatingIPFailStatusMocked test mocked function
func AttachFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: floatingIPIn.AttachedServerID})
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.AttachFloatingIP(mapIn, floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return serverOut
}

// AttachFloatingIPFailJSONMocked test mocked function
func AttachFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.AttachFloatingIP(mapIn, floatingIPIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// DetachFloatingIPMocked test mocked function
func DetachFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID)).Return(dIn, 200, nil)
	err = ds.DetachFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error detaching floating IP")
}

// DetachFloatingIPFailErrMocked test mocked function
func DetachFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DetachFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DetachFloatingIPFailStatusMocked test mocked function
func DetachFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s/attached_server", floatingIPIn.ID)).Return(dIn, 499, nil)
	err = ds.DetachFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteFloatingIPMocked test mocked function
func DeleteFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error deleting floating IP")
}

// DeleteFloatingIPFailErrMocked test mocked function
func DeleteFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteFloatingIPFailStatusMocked test mocked function
func DeleteFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s", floatingIPIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DiscardFloatingIPMocked test mocked function
func DiscardFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s/discard", floatingIPIn.ID)).Return(dIn, 200, nil)
	err = ds.DiscardFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error discarding floating IP")
}

// DiscardFloatingIPFailErrMocked test mocked function
func DiscardFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s/discard", floatingIPIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DiscardFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DiscardFloatingIPFailStatusMocked test mocked function
func DiscardFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/floating_ips/%s/discard", floatingIPIn.ID)).Return(dIn, 499, nil)
	err = ds.DiscardFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
