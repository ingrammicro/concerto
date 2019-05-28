package network

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSubnetServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSubnetService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetSubnetList(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	GetSubnetListMocked(t, subnetsIn)
	GetSubnetListFailErrMocked(t, subnetsIn)
	GetSubnetListFailStatusMocked(t, subnetsIn)
	GetSubnetListFailJSONMocked(t, subnetsIn)
}

func TestGetSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		GetSubnetMocked(t, subnetIn)
		GetSubnetFailErrMocked(t, subnetIn)
		GetSubnetFailStatusMocked(t, subnetIn)
		GetSubnetFailJSONMocked(t, subnetIn)
	}
}

func TestCreateSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		CreateSubnetMocked(t, subnetIn)
		CreateSubnetFailErrMocked(t, subnetIn)
		CreateSubnetFailStatusMocked(t, subnetIn)
		CreateSubnetFailJSONMocked(t, subnetIn)
	}
}

func TestUpdateSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		UpdateSubnetMocked(t, subnetIn)
		UpdateSubnetFailErrMocked(t, subnetIn)
		UpdateSubnetFailStatusMocked(t, subnetIn)
		UpdateSubnetFailJSONMocked(t, subnetIn)
	}
}

func TestDeleteSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		DeleteSubnetMocked(t, subnetIn)
		DeleteSubnetFailErrMocked(t, subnetIn)
		DeleteSubnetFailStatusMocked(t, subnetIn)
	}
}

func TestGetSubnetServersList(t *testing.T) {
	subnetsIn := testdata.GetSubnetServersData()
	GetSubnetServersListMocked(t, subnetsIn)
	GetSubnetServersListFailErrMocked(t, subnetsIn)
	GetSubnetServersListFailStatusMocked(t, subnetsIn)
	GetSubnetServersListFailJSONMocked(t, subnetsIn)
}
