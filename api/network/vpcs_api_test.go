package network

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVPCServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewVPCService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetVPCList(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	GetVPCListMocked(t, vpcsIn)
	GetVPCListFailErrMocked(t, vpcsIn)
	GetVPCListFailStatusMocked(t, vpcsIn)
	GetVPCListFailJSONMocked(t, vpcsIn)
}

func TestGetVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		GetVPCMocked(t, vpcIn)
		GetVPCFailErrMocked(t, vpcIn)
		GetVPCFailStatusMocked(t, vpcIn)
		GetVPCFailJSONMocked(t, vpcIn)
	}
}

func TestCreateVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		CreateVPCMocked(t, vpcIn)
		CreateVPCFailErrMocked(t, vpcIn)
		CreateVPCFailStatusMocked(t, vpcIn)
		CreateVPCFailJSONMocked(t, vpcIn)
	}
}

func TestUpdateVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		UpdateVPCMocked(t, vpcIn)
		UpdateVPCFailErrMocked(t, vpcIn)
		UpdateVPCFailStatusMocked(t, vpcIn)
		UpdateVPCFailJSONMocked(t, vpcIn)
	}
}

func TestDeleteVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		DeleteVPCMocked(t, vpcIn)
		DeleteVPCFailErrMocked(t, vpcIn)
		DeleteVPCFailStatusMocked(t, vpcIn)
	}
}

func TestDiscardVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		DiscardVPCMocked(t, vpcIn)
		DiscardVPCFailErrMocked(t, vpcIn)
		DiscardVPCFailStatusMocked(t, vpcIn)
	}
}
