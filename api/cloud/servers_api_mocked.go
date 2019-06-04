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

// GetServerListMocked test mocked function
func GetServerListMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", "/cloud/servers").Return(dIn, 200, nil)
	serversOut, err := ds.GetServerList()
	assert.Nil(err, "Error getting server list")
	assert.Equal(serversIn, serversOut, "GetServerList returned different servers")

	return serversOut
}

// GetServerListFailErrMocked test mocked function
func GetServerListFailErrMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", "/cloud/servers").Return(dIn, 200, fmt.Errorf("mocked error"))
	serversOut, err := ds.GetServerList()
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serversOut
}

// GetServerListFailStatusMocked test mocked function
func GetServerListFailStatusMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", "/cloud/servers").Return(dIn, 499, nil)
	serversOut, err := ds.GetServerList()
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serversOut
}

// GetServerListFailJSONMocked test mocked function
func GetServerListFailJSONMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/cloud/servers").Return(dIn, 200, nil)
	serversOut, err := ds.GetServerList()
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serversOut
}

// GetServerMocked test mocked function
func GetServerMocked(t *testing.T, server *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(server)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s", server.ID)).Return(dIn, 200, nil)
	serverOut, err := ds.GetServer(server.ID)
	assert.Nil(err, "Error getting server")
	assert.Equal(*server, *serverOut, "GetServer returned different servers")

	return serverOut
}

// GetServerFailErrMocked test mocked function
func GetServerFailErrMocked(t *testing.T, server *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(server)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s", server.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.GetServer(server.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// GetServerFailStatusMocked test mocked function
func GetServerFailStatusMocked(t *testing.T, server *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(server)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s", server.ID)).Return(dIn, 499, nil)
	serverOut, err := ds.GetServer(server.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// GetServerFailJSONMocked test mocked function
func GetServerFailJSONMocked(t *testing.T, server *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s", server.ID)).Return(dIn, 200, nil)
	serverOut, err := ds.GetServer(server.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// CreateServerMocked test mocked function
func CreateServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", "/cloud/servers/", mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.CreateServer(mapIn)
	assert.Nil(err, "Error creating server list")
	assert.Equal(serverIn, serverOut, "CreateServer returned different servers")

	return serverOut
}

// CreateServerFailErrMocked test mocked function
func CreateServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", "/cloud/servers/", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.CreateServer(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// CreateServerFailStatusMocked test mocked function
func CreateServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", "/cloud/servers/", mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.CreateServer(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// CreateServerFailJSONMocked test mocked function
func CreateServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/cloud/servers/", mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.CreateServer(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// UpdateServerMocked test mocked function
func UpdateServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s", serverIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.UpdateServer(mapIn, serverIn.ID)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "UpdateServer returned different servers")

	return serverOut
}

// UpdateServerFailErrMocked test mocked function
func UpdateServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s", serverIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.UpdateServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// UpdateServerFailStatusMocked test mocked function
func UpdateServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s", serverIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.UpdateServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// UpdateServerFailJSONMocked test mocked function
func UpdateServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s", serverIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.UpdateServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// BootServerMocked test mocked function
func BootServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/boot", serverIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.BootServer(mapIn, serverIn.ID)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "BootServer returned different servers")

	return serverOut
}

// BootServerFailErrMocked test mocked function
func BootServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/boot", serverIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.BootServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// BootServerFailStatusMocked test mocked function
func BootServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/boot", serverIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.BootServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// BootServerFailJSONMocked test mocked function
func BootServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/boot", serverIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.BootServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// RebootServerMocked test mocked function
func RebootServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/reboot", serverIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.RebootServer(mapIn, serverIn.ID)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "RebootServer returned different servers")

	return serverOut
}

// RebootServerFailErrMocked test mocked function
func RebootServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/reboot", serverIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.RebootServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// RebootServerFailStatusMocked test mocked function
func RebootServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/reboot", serverIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.RebootServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// RebootServerFailJSONMocked test mocked function
func RebootServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/reboot", serverIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.RebootServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// ShutdownServerMocked test mocked function
func ShutdownServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/shutdown", serverIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.ID)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "ShutdownServer returned different servers")

	return serverOut
}

// ShutdownServerFailErrMocked test mocked function
func ShutdownServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/shutdown", serverIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// ShutdownServerFailStatusMocked test mocked function
func ShutdownServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/shutdown", serverIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// ShutdownServerFailJSONMocked test mocked function
func ShutdownServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/shutdown", serverIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// OverrideServerMocked test mocked function
func OverrideServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/override", serverIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.OverrideServer(mapIn, serverIn.ID)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "OverrideServer returned different servers")

	return serverOut
}

// OverrideServerFailErrMocked test mocked function
func OverrideServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/override", serverIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.OverrideServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// OverrideServerFailStatusMocked test mocked function
func OverrideServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/override", serverIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.OverrideServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// OverrideServerFailJSONMocked test mocked function
func OverrideServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/override", serverIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.OverrideServer(mapIn, serverIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// DeleteServerMocked test mocked function
func DeleteServerMocked(t *testing.T, serverIn *types.Server) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/cloud/servers/%s", serverIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteServer(serverIn.ID)
	assert.Nil(err, "Error deleting server")
}

// DeleteServerFailErrMocked test mocked function
func DeleteServerFailErrMocked(t *testing.T, serverIn *types.Server) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/cloud/servers/%s", serverIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteServer(serverIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteServerFailStatusMocked test mocked function
func DeleteServerFailStatusMocked(t *testing.T, serverIn *types.Server) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/cloud/servers/%s", serverIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteServer(serverIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// GetServerFloatingIPListMocked test mocked function
func GetServerFloatingIPListMocked(t *testing.T, floatingIPsIn []*types.FloatingIP, serverID string) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	fIn, err := json.Marshal(floatingIPsIn)
	assert.Nil(err, "Server floating IP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/floating_ips", serverID)).Return(fIn, 200, nil)
	floatingIPsOut, err := ds.GetServerFloatingIPList(serverID)
	assert.Nil(err, "Error getting server floating IP list")
	assert.Equal(floatingIPsIn, floatingIPsOut, "GetServerFloatingIPListMocked returned different server floating IPs")

	return floatingIPsOut
}

// GetServerFloatingIPListFailErrMocked test mocked function
func GetServerFloatingIPListFailErrMocked(t *testing.T, floatingIPsIn []*types.FloatingIP, serverID string) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	fIn, err := json.Marshal(floatingIPsIn)
	assert.Nil(err, "Server floating IP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/floating_ips", serverID)).Return(fIn, 200, fmt.Errorf("mocked error"))
	floatingIPsOut, err := ds.GetServerFloatingIPList(serverID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPsOut
}

// GetServerFloatingIPListFailStatusMocked test mocked function
func GetServerFloatingIPListFailStatusMocked(t *testing.T, floatingIPsIn []*types.FloatingIP, serverID string) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	fIn, err := json.Marshal(floatingIPsIn)
	assert.Nil(err, "Server floating IP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/floating_ips", serverID)).Return(fIn, 499, nil)
	floatingIPsOut, err := ds.GetServerFloatingIPList(serverID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPsOut
}

// GetServerFloatingIPListFailJSONMocked test mocked function
func GetServerFloatingIPListFailJSONMocked(t *testing.T, floatingIPsIn []*types.FloatingIP, serverID string) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	fIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/floating_ips", serverID)).Return(fIn, 200, nil)
	floatingIPsOut, err := ds.GetServerFloatingIPList(serverID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPsOut
}

// GetServerVolumesListMocked test mocked function
func GetServerVolumesListMocked(t *testing.T, volumesIn []*types.Volume, serverID string) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	vIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Server volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/volumes", serverID)).Return(vIn, 200, nil)
	vOut, err := ds.GetServerVolumesList(serverID)
	assert.Nil(err, "Error getting server volume list")
	assert.Equal(volumesIn, vOut, "GetServerVolumesListMocked returned different server volumes")

	return vOut
}

// GetServerVolumesListFailErrMocked test mocked function
func GetServerVolumesListFailErrMocked(t *testing.T, volumesIn []*types.Volume, serverID string) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	vIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Server volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/volumes", serverID)).Return(vIn, 200, fmt.Errorf("mocked error"))
	vOut, err := ds.GetServerVolumesList(serverID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vOut
}

// GetServerVolumesListFailStatusMocked test mocked function
func GetServerVolumesListFailStatusMocked(t *testing.T, volumesIn []*types.Volume, serverID string) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	vIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Server volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/volumes", serverID)).Return(vIn, 499, nil)
	vOut, err := ds.GetServerVolumesList(serverID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vOut
}

// GetServerVolumesListFailJSONMocked test mocked function
func GetServerVolumesListFailJSONMocked(t *testing.T, volumesIn []*types.Volume, serverID string) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	vIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/volumes", serverID)).Return(vIn, 200, nil)
	vOut, err := ds.GetServerVolumesList(serverID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vOut
}

// GetServerEventListMocked test mocked function
func GetServerEventListMocked(t *testing.T, eventsIn []*types.Event, serverID string) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	evIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Server event test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/events", serverID)).Return(evIn, 200, nil)
	evOut, err := ds.GetEventsList(serverID)
	assert.Nil(err, "Error getting server event list")
	assert.Equal(eventsIn, evOut, "GetServerEventList returned different server events")

	return evOut
}

// GetServerEventListFailErrMocked test mocked function
func GetServerEventListFailErrMocked(t *testing.T, eventsIn []*types.Event, serverID string) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	evIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Server event test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/events", serverID)).Return(evIn, 200, fmt.Errorf("mocked error"))
	evOut, err := ds.GetEventsList(serverID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(evOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return evOut
}

// GetServerEventListFailStatusMocked test mocked function
func GetServerEventListFailStatusMocked(t *testing.T, eventsIn []*types.Event, serverID string) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	evIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Server event test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/events", serverID)).Return(evIn, 499, nil)
	evOut, err := ds.GetEventsList(serverID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(evOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return evOut
}

// GetServerEventListFailJSONMocked test mocked function
func GetServerEventListFailJSONMocked(t *testing.T, eventsIn []*types.Event, serverID string) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	evIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/events", serverID)).Return(evIn, 200, nil)
	evOut, err := ds.GetEventsList(serverID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(evOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return evOut
}

// GetOperationalScriptListMocked test mocked function
func GetOperationalScriptListMocked(t *testing.T, scriptsIn []*types.ScriptChar, serverID string) []*types.ScriptChar {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	oscIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 200, nil)
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)
	assert.Nil(err, "Error getting operational script list")
	assert.Equal(scriptsIn, scriptsOut, "GetOperationalScriptList returned different operational scripts")

	return scriptsOut
}

// GetOperationalScriptFailErrMocked test mocked function
func GetOperationalScriptFailErrMocked(t *testing.T, scriptsIn []*types.ScriptChar, serverID string) []*types.ScriptChar {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	oscIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 200, fmt.Errorf("mocked error"))
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scriptsOut
}

// GetOperationalScriptFailStatusMocked test mocked function
func GetOperationalScriptFailStatusMocked(t *testing.T, scriptsIn []*types.ScriptChar, serverID string) []*types.ScriptChar {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	oscIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 499, nil)
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptsOut
}

// GetOperationalScriptFailJSONMocked test mocked function
func GetOperationalScriptFailJSONMocked(t *testing.T, scriptsIn []*types.ScriptChar, serverID string) []*types.ScriptChar {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	oscIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 200, nil)
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptsOut
}

// ExecuteOperationalScriptListMocked test mocked function
func ExecuteOperationalScriptListMocked(t *testing.T, scriptIn *types.ScriptChar, serverID string, eventDataIn *types.Event) *types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	params, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Server operational scripts test data corrupted")
	oscIn, err := json.Marshal(eventDataIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/operational_scripts/%s/execute", serverID, scriptIn.ID), params).Return(oscIn, 200, nil)
	eventDataOut, err := ds.ExecuteOperationalScript(params, serverID, scriptIn.ID)

	assert.Nil(err, "Error executing operational script")
	assert.Equal(eventDataIn, eventDataOut, "ExecuteOperationalScriptList returned different outputs")

	return eventDataOut
}

// ExecuteOperationalScriptFailErrMocked test mocked function
func ExecuteOperationalScriptFailErrMocked(t *testing.T, scriptIn *types.ScriptChar, serverID string) *types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	params, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Server operational scripts test data corrupted")
	oscIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/operational_scripts/%s/execute", serverID, scriptIn.ID), params).Return(oscIn, 200, fmt.Errorf("mocked error"))
	scriptResponseOut, err := ds.ExecuteOperationalScript(params, serverID, scriptIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptResponseOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scriptResponseOut
}

// ExecuteOperationalScriptFailStatusMocked test mocked function
func ExecuteOperationalScriptFailStatusMocked(t *testing.T, scriptIn *types.ScriptChar, serverID string) *types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	params, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Server operational scripts test data corrupted")
	oscIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/operational_scripts/%s/execute", serverID, scriptIn.ID), params).Return(oscIn, 499, nil)
	scriptResponseOut, err := ds.ExecuteOperationalScript(params, serverID, scriptIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptResponseOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptResponseOut
}

// ExecuteOperationalScriptFailJSONMocked test mocked function
func ExecuteOperationalScriptFailJSONMocked(t *testing.T, scriptIn *types.ScriptChar, serverID string) *types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	params, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// wrong json
	oscIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/cloud/servers/%s/operational_scripts/%s/execute", serverID, scriptIn.ID), params).Return(oscIn, 200, nil)
	scriptResponseOut, err := ds.ExecuteOperationalScript(params, serverID, scriptIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptResponseOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptResponseOut
}
