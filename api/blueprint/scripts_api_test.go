package blueprint

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScriptServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewScriptService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetScriptList(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	GetScriptListMocked(t, scriptsIn)
	GetScriptListFailErrMocked(t, scriptsIn)
	GetScriptListFailStatusMocked(t, scriptsIn)
	GetScriptListFailJSONMocked(t, scriptsIn)
}

func TestGetScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range scriptsIn {
		GetScriptMocked(t, scriptIn)
		GetScriptFailErrMocked(t, scriptIn)
		GetScriptFailStatusMocked(t, scriptIn)
		GetScriptFailJSONMocked(t, scriptIn)
	}
}

func TestCreateScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range scriptsIn {
		CreateScriptMocked(t, scriptIn)
		CreateScriptFailErrMocked(t, scriptIn)
		CreateScriptFailStatusMocked(t, scriptIn)
		CreateScriptFailJSONMocked(t, scriptIn)
	}
}

func TestUpdateScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range scriptsIn {
		UpdateScriptMocked(t, scriptIn)
		UpdateScriptFailErrMocked(t, scriptIn)
		UpdateScriptFailStatusMocked(t, scriptIn)
		UpdateScriptFailJSONMocked(t, scriptIn)
	}
}

func TestDeleteScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range scriptsIn {
		DeleteScriptMocked(t, scriptIn)
		DeleteScriptFailErrMocked(t, scriptIn)
		DeleteScriptFailStatusMocked(t, scriptIn)
	}
}

func TestAddScriptAttachment(t *testing.T) {
	attachmentsIn := testdata.GetAttachmentData()
	scriptsIn := testdata.GetScriptData()
	for _, attachmentIn := range attachmentsIn {
		AddScriptAttachmentMocked(t, attachmentIn, scriptsIn[0].ID)
		AddScriptAttachmentFailErrMocked(t, attachmentIn, scriptsIn[0].ID)
		AddScriptAttachmentFailStatusMocked(t, attachmentIn, scriptsIn[0].ID)
		AddScriptAttachmentFailJSONMocked(t, attachmentIn, scriptsIn[0].ID)
	}
}

func TestUploadScriptAttachment(t *testing.T) {
	attachmentsIn := testdata.GetAttachmentData()
	for _, attachmentIn := range attachmentsIn {
		UploadScriptAttachmentMocked(t, attachmentIn)
		UploadScriptAttachmentFailStatusMocked(t, attachmentIn)
		UploadScriptAttachmentFailErrMocked(t, attachmentIn)
	}
}

func TestUploadedScriptAttachment(t *testing.T) {
	attachmentsIn := testdata.GetAttachmentData()
	for _, attachmentIn := range attachmentsIn {
		UploadedScriptAttachmentMocked(t, attachmentIn)
		UploadedScriptAttachmentFailErrMocked(t, attachmentIn)
		UploadedScriptAttachmentFailStatusMocked(t, attachmentIn)
		UploadedScriptAttachmentFailJSONMocked(t, attachmentIn)
	}
}

func TestListScriptAttachments(t *testing.T) {
	attachmentsIn := testdata.GetAttachmentData()
	scriptsIn := testdata.GetScriptData()
	ListScriptAttachmentsMocked(t, attachmentsIn, scriptsIn[0].ID)
	ListScriptAttachmentsFailErrMocked(t, attachmentsIn, scriptsIn[0].ID)
	ListScriptAttachmentsFailStatusMocked(t, attachmentsIn, scriptsIn[0].ID)
	ListScriptAttachmentsFailJSONMocked(t, attachmentsIn, scriptsIn[0].ID)
}
