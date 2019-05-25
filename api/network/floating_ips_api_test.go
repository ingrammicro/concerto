package network

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFloatingIPServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewFloatingIPService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetFloatingIPList(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	GetFloatingIPListMocked(t, floatingIPsIn)
	GetFloatingIPListMockedFilteredByServer(t, floatingIPsIn)
	GetFloatingIPListFailErrMocked(t, floatingIPsIn)
	GetFloatingIPListFailStatusMocked(t, floatingIPsIn)
	GetFloatingIPListFailJSONMocked(t, floatingIPsIn)
}

func TestGetFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		GetFloatingIPMocked(t, floatingIPIn)
		GetFloatingIPFailErrMocked(t, floatingIPIn)
		GetFloatingIPFailStatusMocked(t, floatingIPIn)
		GetFloatingIPFailJSONMocked(t, floatingIPIn)
	}
}

func TestCreateFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		CreateFloatingIPMocked(t, floatingIPIn)
		CreateFloatingIPFailErrMocked(t, floatingIPIn)
		CreateFloatingIPFailStatusMocked(t, floatingIPIn)
		CreateFloatingIPFailJSONMocked(t, floatingIPIn)
	}
}

func TestUpdateFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		UpdateFloatingIPMocked(t, floatingIPIn)
		UpdateFloatingIPFailErrMocked(t, floatingIPIn)
		UpdateFloatingIPFailStatusMocked(t, floatingIPIn)
		UpdateFloatingIPFailJSONMocked(t, floatingIPIn)
	}
}

func TestAttachFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		AttachFloatingIPMocked(t, floatingIPIn)
		AttachFloatingIPFailErrMocked(t, floatingIPIn)
		AttachFloatingIPFailStatusMocked(t, floatingIPIn)
		AttachFloatingIPFailJSONMocked(t, floatingIPIn)
	}
}

func TestDetachFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		DetachFloatingIPMocked(t, floatingIPIn)
		DetachFloatingIPFailErrMocked(t, floatingIPIn)
		DetachFloatingIPFailStatusMocked(t, floatingIPIn)
	}
}

func TestDeleteFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		DeleteFloatingIPMocked(t, floatingIPIn)
		DeleteFloatingIPFailErrMocked(t, floatingIPIn)
		DeleteFloatingIPFailStatusMocked(t, floatingIPIn)
	}
}

func TestDiscardFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	for _, floatingIPIn := range floatingIPsIn {
		DiscardFloatingIPMocked(t, floatingIPIn)
		DiscardFloatingIPFailErrMocked(t, floatingIPIn)
		DiscardFloatingIPFailStatusMocked(t, floatingIPIn)
	}
}
