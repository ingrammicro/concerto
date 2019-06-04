package storage

import (
	"testing"

	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewServerServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewStoragePlanService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetStoragePlan(t *testing.T) {
	storagePlansIn := testdata.GetStoragePlanData()
	for _, storagePlanIn := range storagePlansIn {
		GetStoragePlanMocked(t, storagePlanIn)
		GetStoragePlanFailErrMocked(t, storagePlanIn)
		GetStoragePlanFailStatusMocked(t, storagePlanIn)
		GetStoragePlanFailJSONMocked(t, storagePlanIn)
	}
}
