package storage

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetStoragePlanMocked test mocked function
func GetStoragePlanMocked(t *testing.T, storagePlan *types.StoragePlan) *types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewStoragePlanService(cs)
	assert.Nil(err, "Couldn't load storage plan service")
	assert.NotNil(ds, "Storage plan service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlan)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/storage/plans/%s", storagePlan.ID)).Return(dIn, 200, nil)
	storagePlanOut, err := ds.GetStoragePlan(storagePlan.ID)
	assert.Nil(err, "Error getting storage plan")
	assert.Equal(*storagePlan, *storagePlanOut, "GetStoragePlan returned different storage plans")

	return storagePlanOut
}

// GetStoragePlanFailErrMocked test mocked function
func GetStoragePlanFailErrMocked(t *testing.T, storagePlan *types.StoragePlan) *types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewStoragePlanService(cs)
	assert.Nil(err, "Couldn't load storage plan service")
	assert.NotNil(ds, "Storage plan service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlan)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/storage/plans/%s", storagePlan.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	storagePlanOut, err := ds.GetStoragePlan(storagePlan.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(storagePlanOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return storagePlanOut
}

// GetStoragePlanFailStatusMocked test mocked function
func GetStoragePlanFailStatusMocked(t *testing.T, storagePlan *types.StoragePlan) *types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewStoragePlanService(cs)
	assert.Nil(err, "Couldn't load storage plan service")
	assert.NotNil(ds, "Storage plan service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlan)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/storage/plans/%s", storagePlan.ID)).Return(dIn, 499, nil)
	storagePlanOut, err := ds.GetStoragePlan(storagePlan.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(storagePlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return storagePlanOut
}

// GetStoragePlanFailJSONMocked test mocked function
func GetStoragePlanFailJSONMocked(t *testing.T, storagePlan *types.StoragePlan) *types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewStoragePlanService(cs)
	assert.Nil(err, "Couldn't load storage plan service")
	assert.NotNil(ds, "Storage plan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/storage/plans/%s", storagePlan.ID)).Return(dIn, 200, nil)
	storagePlanOut, err := ds.GetStoragePlan(storagePlan.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(storagePlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return storagePlanOut
}
