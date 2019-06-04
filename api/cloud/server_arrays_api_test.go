package cloud

import (
	"testing"

	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewServerArrayServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServerArrayService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetServerArrayList(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	GetServerArrayListMocked(t, serverArraysIn)
	GetServerArrayListFailErrMocked(t, serverArraysIn)
	GetServerArrayListFailStatusMocked(t, serverArraysIn)
	GetServerArrayListFailJSONMocked(t, serverArraysIn)
}

func TestGetServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		GetServerArrayMocked(t, serverArrayIn)
		GetServerArrayFailErrMocked(t, serverArrayIn)
		GetServerArrayFailStatusMocked(t, serverArrayIn)
		GetServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestCreateServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		CreateServerArrayMocked(t, serverArrayIn)
		CreateServerArrayFailErrMocked(t, serverArrayIn)
		CreateServerArrayFailStatusMocked(t, serverArrayIn)
		CreateServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestUpdateServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		UpdateServerArrayMocked(t, serverArrayIn)
		UpdateServerArrayFailErrMocked(t, serverArrayIn)
		UpdateServerArrayFailStatusMocked(t, serverArrayIn)
		UpdateServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestBootServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		BootServerArrayMocked(t, serverArrayIn)
		BootServerArrayFailErrMocked(t, serverArrayIn)
		BootServerArrayFailStatusMocked(t, serverArrayIn)
		BootServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestShutdownServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		ShutdownServerArrayMocked(t, serverArrayIn)
		ShutdownServerArrayFailErrMocked(t, serverArrayIn)
		ShutdownServerArrayFailStatusMocked(t, serverArrayIn)
		ShutdownServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestEmptyServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		EmptyServerArrayMocked(t, serverArrayIn)
		EmptyServerArrayFailErrMocked(t, serverArrayIn)
		EmptyServerArrayFailStatusMocked(t, serverArrayIn)
		EmptyServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestEnlargeServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		EnlargeServerArrayMocked(t, serverArrayIn)
		EnlargeServerArrayFailErrMocked(t, serverArrayIn)
		EnlargeServerArrayFailStatusMocked(t, serverArrayIn)
		EnlargeServerArrayFailJSONMocked(t, serverArrayIn)
	}
}

func TestGetServerArrayServerList(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	serversIn := testdata.GetServerData()
	for _, serverArrayIn := range serverArraysIn {
		GetServerArrayServerListMocked(t, serversIn, serverArrayIn.ID)
		GetServerArrayServerListFailErrMocked(t, serversIn, serverArrayIn.ID)
		GetServerArrayServerListFailStatusMocked(t, serversIn, serverArrayIn.ID)
		GetServerArrayServerListFailJSONMocked(t, serversIn, serverArrayIn.ID)
	}
}

func TestDeleteServerArray(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	for _, serverArrayIn := range serverArraysIn {
		DeleteServerArrayMocked(t, serverArrayIn)
		DeleteServerArrayFailErrMocked(t, serverArrayIn)
		DeleteServerArrayFailStatusMocked(t, serverArrayIn)
	}
}
