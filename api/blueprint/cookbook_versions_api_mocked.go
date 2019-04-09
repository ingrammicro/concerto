package blueprint

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetCookbookVersionListMocked test mocked function
func GetCookbookVersionListMocked(t *testing.T, cbsIn *[]types.CookbookVersion) *[]types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbsIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Get", "/v2/blueprint/cookbook_versions").Return(dIn, 200, nil)
	cbsOut, err := ds.GetCookbookVersionList()
	assert.Nil(err, "Error getting cookbook version list")
	assert.Equal(*cbsIn, cbsOut, "GetCookbookVersionList returned different services")

	return &cbsOut
}

// GetCookbookVersionListFailErrMocked test mocked function
func GetCookbookVersionListFailErrMocked(t *testing.T, cbsIn *[]types.CookbookVersion) *[]types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbsIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Get", "/v2/blueprint/cookbook_versions").Return(dIn, 200, fmt.Errorf("Mocked error"))
	cbsOut, err := ds.GetCookbookVersionList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cbsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &cbsOut
}

// GetCookbookVersionListFailStatusMocked test mocked function
func GetCookbookVersionListFailStatusMocked(t *testing.T, cbsIn *[]types.CookbookVersion) *[]types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbsIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Get", "/v2/blueprint/cookbook_versions").Return(dIn, 499, nil)
	cbsOut, err := ds.GetCookbookVersionList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cbsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &cbsOut
}

// GetCookbookVersionListFailJSONMocked test mocked function
func GetCookbookVersionListFailJSONMocked(t *testing.T, cbsIn *[]types.CookbookVersion) *[]types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v2/blueprint/cookbook_versions").Return(dIn, 200, nil)
	cbsOut, err := ds.GetCookbookVersionList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cbsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &cbsOut
}

// GetCookbookVersionMocked test mocked function
func GetCookbookVersionMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 200, nil)
	cbOut, err := ds.GetCookbookVersion(cbIn.ID)
	assert.Nil(err, "Error getting cookbook version")
	assert.Equal(*cbIn, *cbOut, "GetCookbookVersion returned different services")

	return cbOut
}

// GetCookbookVersionFailErrMocked test mocked function
func GetCookbookVersionFailErrMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	cbOut, err := ds.GetCookbookVersion(cbIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return cbOut
}

// GetCookbookVersionFailStatusMocked test mocked function
func GetCookbookVersionFailStatusMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 499, nil)
	cbOut, err := ds.GetCookbookVersion(cbIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cbOut
}

// GetCookbookVersionFailJSONMocked test mocked function
func GetCookbookVersionFailJSONMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 200, nil)
	cbOut, err := ds.GetCookbookVersion(cbIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cbOut
}

// CreateCookbookVersionMocked test mocked function
func CreateCookbookVersionMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Post", "/v2/blueprint/cookbook_versions", mapIn).Return(dOut, 201, nil)
	cbOut, err := ds.CreateCookbookVersion(mapIn)
	assert.Nil(err, "Error creating cookbook version")
	assert.Equal(cbIn, cbOut, "CreateCookbookVersion returned different cookbook version")

	return cbOut
}

// CreateCookbookVersionFailErrMocked test mocked function
func CreateCookbookVersionFailErrMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Post", "/v2/blueprint/cookbook_versions", mapIn).Return(dOut, 201, fmt.Errorf("Mocked error"))
	cbOut, err := ds.CreateCookbookVersion(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return cbOut
}

// CreateCookbookVersionFailStatusMocked test mocked function
func CreateCookbookVersionFailStatusMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Post", "/v2/blueprint/cookbook_versions", mapIn).Return(dOut, 499, nil)
	cbOut, err := ds.CreateCookbookVersion(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cbOut
}

// CreateCookbookVersionFailJSONMocked test mocked function
func CreateCookbookVersionFailJSONMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v2/blueprint/cookbook_versions", mapIn).Return(dIn, 201, nil)
	cbOut, err := ds.CreateCookbookVersion(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cbOut
}

// UploadCookbookVersionMocked test mocked function
func UploadCookbookVersionMocked(t *testing.T, cbIn *types.CookbookVersion) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := cbIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 200, nil)
	err = ds.UploadCookbookVersion(sourceFilePath, targetURL)
	assert.Nil(err, "Error uploading cookbook version file")
}

// UploadCookbookVersionFailStatusMocked test mocked function
func UploadCookbookVersionFailStatusMocked(t *testing.T, cbIn *types.CookbookVersion) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := cbIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, nil)
	err = ds.UploadCookbookVersion(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
}

// UploadCookbookVersionFailErrMocked test mocked function
func UploadCookbookVersionFailErrMocked(t *testing.T, cbIn *types.CookbookVersion) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := cbIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, fmt.Errorf("Mocked error"))
	err = ds.UploadCookbookVersion(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// ProcessCookbookVersionMocked test mocked function
func ProcessCookbookVersionMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s/process", cbIn.ID), mapIn).Return(dOut, 200, nil)
	cbOut, err := ds.ProcessCookbookVersion(mapIn, cbIn.ID)
	assert.Nil(err, "Error processing cookbook version")
	assert.Equal(cbIn, cbOut, "ProcessCookbookVersion returned different cookbook version")

	return cbOut
}

// ProcessCookbookVersionFailErrMocked test mocked function
func ProcessCookbookVersionFailErrMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s/process", cbIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	cbOut, err := ds.ProcessCookbookVersion(mapIn, cbIn.ID)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return cbOut
}

// ProcessCookbookVersionFailStatusMocked test mocked function
func ProcessCookbookVersionFailStatusMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s/process", cbIn.ID), mapIn).Return(dOut, 499, nil)
	cbOut, err := ds.ProcessCookbookVersion(mapIn, cbIn.ID)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cbOut
}

// ProcessCookbookVersionFailJSONMocked test mocked function
func ProcessCookbookVersionFailJSONMocked(t *testing.T, cbIn *types.CookbookVersion) *types.CookbookVersion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s/process", cbIn.ID), mapIn).Return(dIn, 201, nil)
	cbOut, err := ds.ProcessCookbookVersion(mapIn, cbIn.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cbOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cbOut
}

// DeleteCookbookVersionMocked test mocked function
func DeleteCookbookVersionMocked(t *testing.T, cbIn *types.CookbookVersion) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteCookbookVersion(cbIn.ID)
	assert.Nil(err, "Error deleting cookbook version")

}

// DeleteCookbookVersionFailErrMocked test mocked function
func DeleteCookbookVersionFailErrMocked(t *testing.T, cbIn *types.CookbookVersion) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteCookbookVersion(cbIn.ID)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// DeleteCookbookVersionFailStatusMocked test mocked function
func DeleteCookbookVersionFailStatusMocked(t *testing.T, cbIn *types.CookbookVersion) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCookbookVersionService(cs)
	assert.Nil(err, "Couldn't load cookbook version service")
	assert.NotNil(ds, "CookbookVersion service not instanced")

	// to json
	dIn, err := json.Marshal(cbIn)
	assert.Nil(err, "CookbookVersion test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v2/blueprint/cookbook_versions/%s", cbIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteCookbookVersion(cbIn.ID)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

}
