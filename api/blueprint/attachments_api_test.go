package blueprint

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAttachmentServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewAttachmentService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetAttachment(t *testing.T) {
	attachmentsIn := testdata.GetAttachmentData()
	for _, attachmentIn := range attachmentsIn {
		GetAttachmentMocked(t, attachmentIn)
		GetAttachmentFailErrMocked(t, attachmentIn)
		GetAttachmentFailStatusMocked(t, attachmentIn)
		GetAttachmentFailJSONMocked(t, attachmentIn)
	}
}

func TestDownloadAttachment(t *testing.T) {
	dataIn := testdata.GetDownloadAttachmentData()
	DownloadAttachmentMocked(t, dataIn)
	DownloadAttachmentFailErrMocked(t, dataIn)
}

func TestDeleteAttachment(t *testing.T) {
	attachmentsIn := testdata.GetAttachmentData()
	for _, attachmentIn := range attachmentsIn {
		DeleteAttachmentMocked(t, attachmentIn)
		DeleteAttachmentFailErrMocked(t, attachmentIn)
		DeleteAttachmentFailStatusMocked(t, attachmentIn)
	}
}
