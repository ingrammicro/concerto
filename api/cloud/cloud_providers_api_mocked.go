package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetCloudProviderListMocked test mocked function
func GetCloudProviderListMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.GetCloudProviderList()
	assert.Nil(err, "Error getting cloudProvider list")
	assert.Equal(cloudProvidersIn, cloudProvidersOut, "GetCloudProviderList returned different cloudProviders")

	return cloudProvidersOut
}

// GetCloudProviderListFailErrMocked test mocked function
func GetCloudProviderListFailErrMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudProvidersOut, err := ds.GetCloudProviderList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudProvidersOut
}

// GetCloudProviderListFailStatusMocked test mocked function
func GetCloudProviderListFailStatusMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 499, nil)
	cloudProvidersOut, err := ds.GetCloudProviderList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudProvidersOut
}

// GetCloudProviderListFailJSONMocked test mocked function
func GetCloudProviderListFailJSONMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.GetCloudProviderList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudProvidersOut
}

// GetServerStoragePlanListMocked test mocked function
func GetServerStoragePlanListMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlansIn)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 200, nil)
	storagePlansOut, err := ds.GetServerStoragePlanList(providerID)
	assert.Nil(err, "Error getting storage plan list")
	assert.Equal(storagePlansIn, storagePlansOut, "GetServerStoragePlanList returned different storage plans")

	return storagePlansOut
}

// GetServerStoragePlanListFailErrMocked test mocked function
func GetServerStoragePlanListFailErrMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlansIn)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	storagePlansOut, err := ds.GetServerStoragePlanList(providerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(storagePlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return storagePlansOut
}

// GetServerStoragePlanListFailStatusMocked test mocked function
func GetServerStoragePlanListFailStatusMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlansIn)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 499, nil)
	storagePlansOut, err := ds.GetServerStoragePlanList(providerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(storagePlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return storagePlansOut
}

// GetServerStoragePlanListFailJSONMocked test mocked function
func GetServerStoragePlanListFailJSONMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 200, nil)
	storagePlansOut, err := ds.GetServerStoragePlanList(providerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(storagePlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return storagePlansOut
}
