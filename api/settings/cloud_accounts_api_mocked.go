package settings

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetCloudAccountListMocked test mocked function
func GetCloudAccountListMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := clAccService.GetCloudAccountList()
	assert.Nil(err, "Error getting cloudAccount list")
	assert.Equal(*cloudAccountsIn, cloudAccountsOut, "GetCloudAccountList returned different cloudAccounts")

	return &cloudAccountsOut
}

// GetCloudAccountListFailErrMocked test mocked function
func GetCloudAccountListFailErrMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 200, fmt.Errorf("Mocked error"))
	cloudAccountsOut, err := clAccService.GetCloudAccountList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &cloudAccountsOut
}

// GetCloudAccountListFailStatusMocked test mocked function
func GetCloudAccountListFailStatusMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 499, nil)
	cloudAccountsOut, err := clAccService.GetCloudAccountList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &cloudAccountsOut
}

// GetCloudAccountListFailJSONMocked test mocked function
func GetCloudAccountListFailJSONMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := clAccService.GetCloudAccountList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &cloudAccountsOut
}
