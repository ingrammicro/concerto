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

// GetAttachmentMocked test mocked function
func GetAttachmentMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Attachment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 200, nil)
	attachmentOut, err := ds.GetAttachment(attachmentIn.ID)
	assert.Nil(err, "Error getting attachment")
	assert.Equal(*attachmentIn, *attachmentOut, "GetAttachment returned different attachments")

	return attachmentOut
}

// GetAttachmentFailErrMocked test mocked function
func GetAttachmentFailErrMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Attachment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	attachmentOut, err := ds.GetAttachment(attachmentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return attachmentOut
}

// GetAttachmentFailStatusMocked test mocked function
func GetAttachmentFailStatusMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Attachment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 499, nil)
	attachmentOut, err := ds.GetAttachment(attachmentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return attachmentOut
}

// GetAttachmentFailJSONMocked test mocked function
func GetAttachmentFailJSONMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 200, nil)
	attachmentOut, err := ds.GetAttachment(attachmentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return attachmentOut
}

// DownloadAttachmentMocked test mocked function
func DownloadAttachmentMocked(t *testing.T, dataIn map[string]string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

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
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	urlSource := dataIn["fakeEndpoint"]
	pathFile := dataIn["fakeAttachmentDir"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return("", 499, fmt.Errorf("mocked error"))
	_, status, err := ds.DownloadAttachment(urlSource, pathFile)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(status, 499, "DownloadAttachment returned an unexpected status code")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteAttachmentMocked test mocked function
func DeleteAttachmentMocked(t *testing.T, attachmentIn *types.Attachment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Attachment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteAttachment(attachmentIn.ID)
	assert.Nil(err, "Error deleting attachment")

}

// DeleteAttachmentFailErrMocked test mocked function
func DeleteAttachmentFailErrMocked(t *testing.T, attachmentIn *types.Attachment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Attachment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteAttachment(attachmentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteAttachmentFailStatusMocked test mocked function
func DeleteAttachmentFailStatusMocked(t *testing.T, attachmentIn *types.Attachment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAttachmentService(cs)
	assert.Nil(err, "Couldn't load attachment service")
	assert.NotNil(ds, "Attachment service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Attachment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/blueprint/attachments/%s", attachmentIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteAttachment(attachmentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
