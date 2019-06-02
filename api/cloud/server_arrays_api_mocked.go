package cloud

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetServerArrayListMocked test mocked function
func GetServerArrayListMocked(t *testing.T, serverArraysIn []*types.ServerArray) []*types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArraysIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", "/cloud/server_arrays").Return(dIn, 200, nil)
	serverArraysOut, err := ds.GetServerArrayList()
	assert.Nil(err, "Error getting server array list")
	assert.Equal(serverArraysIn, serverArraysOut, "GetServerArrayList returned different server arrays")

	return serverArraysOut
}

// GetServerArrayListFailErrMocked test mocked function
func GetServerArrayListFailErrMocked(t *testing.T, serverArraysIn []*types.ServerArray) []*types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArraysIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", "/cloud/server_arrays").Return(dIn, 200, fmt.Errorf("mocked error"))
	serverArraysOut, err := ds.GetServerArrayList()
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArraysOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArraysOut
}

// GetServerArrayListFailStatusMocked test mocked function
func GetServerArrayListFailStatusMocked(t *testing.T, serverArraysIn []*types.ServerArray) []*types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArraysIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", "/cloud/server_arrays").Return(dIn, 499, nil)
	serverArraysOut, err := ds.GetServerArrayList()
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArraysOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArraysOut
}

// GetServerArrayListFailJSONMocked test mocked function
func GetServerArrayListFailJSONMocked(t *testing.T, serverArraysIn []*types.ServerArray) []*types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/cloud/server_arrays").Return(dIn, 200, nil)
	serverArraysOut, err := ds.GetServerArrayList()
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArraysOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArraysOut
}

// GetServerArrayMocked test mocked function
func GetServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 200, nil)
	serverArrayOut, err := ds.GetServerArray(serverArrayIn.ID)
	assert.Nil(err, "Error getting server array")
	assert.Equal(*serverArrayIn, *serverArrayOut, "GetServerArray returned different server arrays")

	return serverArrayOut
}

// GetServerArrayFailErrMocked test mocked function
func GetServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.GetServerArray(serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// GetServerArrayFailStatusMocked test mocked function
func GetServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 499, nil)
	serverArrayOut, err := ds.GetServerArray(serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// GetServerArrayFailJSONMocked test mocked function
func GetServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 200, nil)
	serverArrayOut, err := ds.GetServerArray(serverArrayIn.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// CreateServerArrayMocked test mocked function
func CreateServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Post", "/cloud/server_arrays/", mapIn).Return(dOut, 200, nil)
	serverArrayOut, err := ds.CreateServerArray(mapIn)
	assert.Nil(err, "Error creating server array")
	assert.Equal(serverArrayIn, serverArrayOut, "CreateServerArray returned different server arrays")

	return serverArrayOut
}

// CreateServerArrayFailErrMocked test mocked function
func CreateServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Post", "/cloud/server_arrays/", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.CreateServerArray(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// CreateServerArrayFailStatusMocked test mocked function
func CreateServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Post", "/cloud/server_arrays/", mapIn).Return(dOut, 499, nil)
	serverArrayOut, err := ds.CreateServerArray(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// CreateServerArrayFailJSONMocked test mocked function
func CreateServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/cloud/server_arrays/", mapIn).Return(dIn, 200, nil)
	serverArrayOut, err := ds.CreateServerArray(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// UpdateServerArrayMocked test mocked function
func UpdateServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID), mapIn).Return(dOut, 200, nil)
	serverArrayOut, err := ds.UpdateServerArray(mapIn, serverArrayIn.ID)
	assert.Nil(err, "Error updating server array")
	assert.Equal(serverArrayIn, serverArrayOut, "UpdateServerArray returned different server arrays")

	return serverArrayOut
}

// UpdateServerArrayFailErrMocked test mocked function
func UpdateServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.UpdateServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// UpdateServerArrayFailStatusMocked test mocked function
func UpdateServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID), mapIn).Return(dOut, 499, nil)
	serverArrayOut, err := ds.UpdateServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// UpdateServerArrayFailJSONMocked test mocked function
func UpdateServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID), mapIn).Return(dIn, 200, nil)
	serverArrayOut, err := ds.UpdateServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// BootServerArrayMocked test mocked function
func BootServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/boot", serverArrayIn.ID), mapIn).Return(dOut, 200, nil)
	serverArrayOut, err := ds.BootServerArray(mapIn, serverArrayIn.ID)
	assert.Nil(err, "Error booting server array")
	assert.Equal(serverArrayIn, serverArrayOut, "BootServerArray returned different server arrays")

	return serverArrayOut
}

// BootServerArrayFailErrMocked test mocked function
func BootServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/boot", serverArrayIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.BootServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// BootServerArrayFailStatusMocked test mocked function
func BootServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/boot", serverArrayIn.ID), mapIn).Return(dOut, 499, nil)
	serverArrayOut, err := ds.BootServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// BootServerArrayFailJSONMocked test mocked function
func BootServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/boot", serverArrayIn.ID), mapIn).Return(dIn, 200, nil)
	serverArrayOut, err := ds.BootServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// ShutdownServerArrayMocked test mocked function
func ShutdownServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/shutdown", serverArrayIn.ID), mapIn).Return(dOut, 200, nil)
	serverArrayOut, err := ds.ShutdownServerArray(mapIn, serverArrayIn.ID)
	assert.Nil(err, "Error shutting down server array")
	assert.Equal(serverArrayIn, serverArrayOut, "ShutdownServerArray returned different server arrays")

	return serverArrayOut
}

// ShutdownServerArrayFailErrMocked test mocked function
func ShutdownServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/shutdown", serverArrayIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.ShutdownServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// ShutdownServerArrayFailStatusMocked test mocked function
func ShutdownServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/shutdown", serverArrayIn.ID), mapIn).Return(dOut, 499, nil)
	serverArrayOut, err := ds.ShutdownServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// ShutdownServerArrayFailJSONMocked test mocked function
func ShutdownServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/shutdown", serverArrayIn.ID), mapIn).Return(dIn, 200, nil)
	serverArrayOut, err := ds.ShutdownServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// EmptyServerArrayMocked test mocked function
func EmptyServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/empty", serverArrayIn.ID), mapIn).Return(dOut, 200, nil)
	serverArrayOut, err := ds.EmptyServerArray(mapIn, serverArrayIn.ID)
	assert.Nil(err, "Error emptying server array")
	assert.Equal(serverArrayIn, serverArrayOut, "EmptyServerArray returned different server arrays")

	return serverArrayOut
}

// EmptyServerArrayFailErrMocked test mocked function
func EmptyServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/empty", serverArrayIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.EmptyServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// EmptyServerArrayFailStatusMocked test mocked function
func EmptyServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/empty", serverArrayIn.ID), mapIn).Return(dOut, 499, nil)
	serverArrayOut, err := ds.EmptyServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// EmptyServerArrayFailJSONMocked test mocked function
func EmptyServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/server_arrays/%s/empty", serverArrayIn.ID), mapIn).Return(dIn, 200, nil)
	serverArrayOut, err := ds.EmptyServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// EnlargeServerArrayMocked test mocked function
func EnlargeServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayIn.ID), mapIn).Return(dOut, 200, nil)
	serverArrayOut, err := ds.EnlargeServerArray(mapIn, serverArrayIn.ID)
	assert.Nil(err, "Error enlarging server array")
	assert.Equal(serverArrayIn, serverArrayOut, "EnlargeServerArray returned different server arrays")

	return serverArrayOut
}

// EnlargeServerArrayFailErrMocked test mocked function
func EnlargeServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverArrayOut, err := ds.EnlargeServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverArrayOut
}

// EnlargeServerArrayFailStatusMocked test mocked function
func EnlargeServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// to json
	dOut, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayIn.ID), mapIn).Return(dOut, 499, nil)
	serverArrayOut, err := ds.EnlargeServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverArrayOut
}

// EnlargeServerArrayFailJSONMocked test mocked function
func EnlargeServerArrayFailJSONMocked(t *testing.T, serverArrayIn *types.ServerArray) *types.ServerArray {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayIn.ID), mapIn).Return(dIn, 200, nil)
	serverArrayOut, err := ds.EnlargeServerArray(mapIn, serverArrayIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverArrayOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverArrayOut
}

// GetServerArrayServerListMocked test mocked function
func GetServerArrayServerListMocked(t *testing.T, serversIn []*types.Server, serverArrayID string) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayID)).Return(dIn, 200, nil)
	serversOut, err := ds.GetServerArrayServerList(serverArrayID)
	assert.Nil(err, "Error getting server list")
	assert.Equal(serversIn, serversOut, "GetServerArrayServerList returned different servers")

	return serversOut
}

// GetServerArrayServerListFailErrMocked test mocked function
func GetServerArrayServerListFailErrMocked(t *testing.T, serversIn []*types.Server, serverArrayID string) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	serversOut, err := ds.GetServerArrayServerList(serverArrayID)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serversOut
}

// GetServerArrayServerListFailStatusMocked test mocked function
func GetServerArrayServerListFailStatusMocked(t *testing.T, serversIn []*types.Server, serverArrayID string) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayID)).Return(dIn, 499, nil)
	serversOut, err := ds.GetServerArrayServerList(serverArrayID)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serversOut
}

// GetServerArrayServerListFailJSONMocked test mocked function
func GetServerArrayServerListFailJSONMocked(t *testing.T, serversIn []*types.Server, serverArrayID string) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/server_arrays/%s/servers", serverArrayID)).Return(dIn, 200, nil)
	serversOut, err := ds.GetServerArrayServerList(serverArrayID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serversOut
}

// DeleteServerArrayMocked test mocked function
func DeleteServerArrayMocked(t *testing.T, serverArrayIn *types.ServerArray) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteServerArray(serverArrayIn.ID)
	assert.Nil(err, "Error deleting server array")
}

// DeleteServerArrayFailErrMocked test mocked function
func DeleteServerArrayFailErrMocked(t *testing.T, serverArrayIn *types.ServerArray) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteServerArray(serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteServerArrayFailStatusMocked test mocked function
func DeleteServerArrayFailStatusMocked(t *testing.T, serverArrayIn *types.ServerArray) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerArrayService(cs)
	assert.Nil(err, "Couldn't load server array service")
	assert.NotNil(ds, "Server array service not instanced")

	// to json
	dIn, err := json.Marshal(serverArrayIn)
	assert.Nil(err, "Server array test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/cloud/server_arrays/%s", serverArrayIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteServerArray(serverArrayIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
