package network

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVPNServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewVPNService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetVPN(t *testing.T) {
	vpcsIn := testdata.GetVPNData()
	for _, vpcIn := range vpcsIn {
		GetVPNMocked(t, vpcIn)
		GetVPNFailErrMocked(t, vpcIn)
		GetVPNFailStatusMocked(t, vpcIn)
		GetVPNFailJSONMocked(t, vpcIn)
	}
}

func TestCreateVPN(t *testing.T) {
	vpcsIn := testdata.GetVPNData()
	for _, vpcIn := range vpcsIn {
		CreateVPNMocked(t, vpcIn)
		CreateVPNFailErrMocked(t, vpcIn)
		CreateVPNFailStatusMocked(t, vpcIn)
		CreateVPNFailJSONMocked(t, vpcIn)
	}
}

func TestDeleteVPN(t *testing.T) {
	vpcsIn := testdata.GetVPNData()
	for _, vpcIn := range vpcsIn {
		DeleteVPNMocked(t, vpcIn)
		DeleteVPNFailErrMocked(t, vpcIn)
		DeleteVPNFailStatusMocked(t, vpcIn)
	}
}

func TestGetVPNListPlans(t *testing.T) {
	vpnPlansIn := testdata.GetVPNPlanData()
	GetVPNListPlansMocked(t, vpnPlansIn, "fakeVpcID0")
	GetVPNListPlansFailErrMocked(t, vpnPlansIn, "fakeVpcID0")
	GetVPNListPlansFailStatusMocked(t, vpnPlansIn, "fakeVpcID0")
	GetVPNListPlansFailJSONMocked(t, vpnPlansIn,"fakeVpcID0")
}
