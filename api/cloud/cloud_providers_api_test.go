package cloud

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCloudProviderServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudProviderService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCloudProviderList(t *testing.T) {
	cloudProvidersIn := testdata.GetCloudProviderData()
	GetCloudProviderListMocked(t, cloudProvidersIn)
	GetCloudProviderListFailErrMocked(t, cloudProvidersIn)
	GetCloudProviderListFailStatusMocked(t, cloudProvidersIn)
	GetCloudProviderListFailJSONMocked(t, cloudProvidersIn)
}

func TestGetServerStoragePlanList(t *testing.T) {
	storagePlansIn := testdata.GetStoragePlanData()
	GetServerStoragePlanListMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	GetServerStoragePlanListFailErrMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	GetServerStoragePlanListFailStatusMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
	GetServerStoragePlanListFailJSONMocked(t, storagePlansIn, storagePlansIn[0].CloudProviderID)
}
