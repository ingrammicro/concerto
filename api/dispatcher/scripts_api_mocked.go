package dispatcher

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetDispatcherScriptCharacterizationsByTypeMocked test mocked function
func GetDispatcherScriptCharacterizationsByTypeMocked(t *testing.T, phase string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// to json
	dIn, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations?type=%s", phase)).Return(dIn, 200, nil)
	scOut, err := ds.GetDispatcherScriptCharacterizationsByType(phase)
	assert.Nil(err, "Error getting dispatcher")
	assert.Equal(scIn, scOut, "GetDispatcherScriptCharacterizationsByType returned different services")

	return scOut
}

// GetDispatcherScriptCharacterizationsByTypeFailErrMocked test mocked function
func GetDispatcherScriptCharacterizationsByTypeFailErrMocked(t *testing.T, phase string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// to json
	dIn, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations?type=%s", phase)).Return(dIn, 200, fmt.Errorf("mocked error"))
	scOut, err := ds.GetDispatcherScriptCharacterizationsByType(phase)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scOut
}

// GetDispatcherScriptCharacterizationsByTypeFailStatusMocked test mocked function
func GetDispatcherScriptCharacterizationsByTypeFailStatusMocked(t *testing.T, phase string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// to json
	dIn, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations?type=%s", phase)).Return(dIn, 499, nil)
	scOut, err := ds.GetDispatcherScriptCharacterizationsByType(phase)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scOut
}

// GetDispatcherScriptCharacterizationsByTypeFailJSONMocked test mocked function
func GetDispatcherScriptCharacterizationsByTypeFailJSONMocked(t *testing.T, phase string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations?type=%s", phase)).Return(dIn, 200, nil)
	scOut, err := ds.GetDispatcherScriptCharacterizationsByType(phase)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scOut
}

// GetDispatcherScriptCharacterizationsByUUIDMocked test mocked function
func GetDispatcherScriptCharacterizationsByUUIDMocked(t *testing.T, scriptCharacterizationUUID string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// to json
	dIn, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations/%s", scriptCharacterizationUUID)).Return(dIn, 200, nil)
	scOut, err := ds.GetDispatcherScriptCharacterizationsByUUID(scriptCharacterizationUUID)
	assert.Nil(err, "Error getting dispatcher")
	assert.Equal(scIn, scOut, "GetDispatcherScriptCharacterizationsByUUID returned different services")

	return scOut
}

// GetDispatcherScriptCharacterizationsByUUIDFailErrMocked test mocked function
func GetDispatcherScriptCharacterizationsByUUIDFailErrMocked(t *testing.T, scriptCharacterizationUUID string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// to json
	dIn, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations/%s", scriptCharacterizationUUID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	scOut, err := ds.GetDispatcherScriptCharacterizationsByUUID(scriptCharacterizationUUID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scOut
}

// GetDispatcherScriptCharacterizationsByUUIDFailStatusMocked test mocked function
func GetDispatcherScriptCharacterizationsByUUIDFailStatusMocked(t *testing.T, scriptCharacterizationUUID string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// to json
	dIn, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations/%s", scriptCharacterizationUUID)).Return(dIn, 499, nil)
	scOut, err := ds.GetDispatcherScriptCharacterizationsByUUID(scriptCharacterizationUUID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scOut
}

// GetDispatcherScriptCharacterizationsByUUIDFailJSONMocked test mocked function
func GetDispatcherScriptCharacterizationsByUUIDFailJSONMocked(t *testing.T, scriptCharacterizationUUID string, scIn []*types.ScriptCharacterization) []*types.ScriptCharacterization {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/script_characterizations/%s", scriptCharacterizationUUID)).Return(dIn, 200, nil)
	scOut, err := ds.GetDispatcherScriptCharacterizationsByUUID(scriptCharacterizationUUID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scOut
}

// ReportScriptConclusionsMocked test mocked function
func ReportScriptConclusionsMocked(t *testing.T, scIn *types.ScriptConclusion) *types.ScriptConclusion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// to json
	dOut, err := json.Marshal(scIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Post", "/blueprint/script_conclusions", mapIn).Return(dOut, 200, nil)
	scOut, _, err := ds.ReportScriptConclusions(mapIn)
	assert.Nil(err, "Error processing dispatcher")
	assert.Equal(scIn, scOut, "ReportScriptConclusions returned different dispatcher")

	return scOut
}

// ReportScriptConclusionsFailErrMocked test mocked function
func ReportScriptConclusionsFailErrMocked(t *testing.T, cbIn *types.ScriptConclusion) *types.ScriptConclusion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Post", "/blueprint/script_conclusions", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	scOut, _, err := ds.ReportScriptConclusions(mapIn)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scOut
}

// ReportScriptConclusionsFailStatusMocked test mocked function
func ReportScriptConclusionsFailStatusMocked(t *testing.T, cbIn *types.ScriptConclusion) *types.ScriptConclusion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// to json
	dOut, err := json.Marshal(cbIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// call service
	cs.On("Post", "/blueprint/script_conclusions", mapIn).Return(dOut, 499, nil)
	scOut, _, err := ds.ReportScriptConclusions(mapIn)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scOut
}

// ReportScriptConclusionsFailJSONMocked test mocked function
func ReportScriptConclusionsFailJSONMocked(t *testing.T, cbIn *types.ScriptConclusion) *types.ScriptConclusion {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cbIn)
	assert.Nil(err, "Dispatcher test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/blueprint/script_conclusions", mapIn).Return(dIn, 201, nil)
	scOut, _, err := ds.ReportScriptConclusions(mapIn)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scOut
}

// DownloadAttachmentMocked test mocked function
func DownloadAttachmentMocked(t *testing.T, dataIn map[string]string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	urlSource := dataIn["fakeEndpoint"]
	pathFile := dataIn["fakeAttachmentDir"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return(pathFile, 200, nil)
	realFileName, status, err := ds.DownloadAttachment(urlSource, pathFile)
	assert.Nil(err, "Error downloading attachment file")
	assert.Equal(status, 200, "DownloadAttachment returned invalid response")
	assert.Equal(realFileName, pathFile, "Invalid downloaded file path")
}

// DownloadAttachmentFailErrMocked test mocked function
func DownloadAttachmentFailErrMocked(t *testing.T, dataIn map[string]string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDispatcherService(cs)
	assert.Nil(err, "Couldn't load dispatcher service")
	assert.NotNil(ds, "Dispatcher service not instanced")

	urlSource := dataIn["fakeEndpoint"]
	pathFile := dataIn["fakeAttachmentDir"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return("", 499, fmt.Errorf("mocked error"))
	_, status, err := ds.DownloadAttachment(urlSource, pathFile)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(status, 499, "DownloadAttachment returned an unexpected status code")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}
