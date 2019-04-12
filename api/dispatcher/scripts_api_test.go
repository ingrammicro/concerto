package dispatcher

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDispatcherServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewDispatcherService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetDispatcherScriptCharacterizationsByType(t *testing.T) {
	dIn := testdata.GetScriptCharacterizationsData()
	GetDispatcherScriptCharacterizationsByTypeMocked(t, "boot", dIn)
	GetDispatcherScriptCharacterizationsByTypeFailErrMocked(t, "boot", dIn)
	GetDispatcherScriptCharacterizationsByTypeFailStatusMocked(t, "boot", dIn)
	GetDispatcherScriptCharacterizationsByTypeFailJSONMocked(t, "boot", dIn)
}

func TestGetDispatcherScriptCharacterizationsByUUID(t *testing.T) {
	dIn := testdata.GetScriptCharacterizationsData()
	GetDispatcherScriptCharacterizationsByUUIDMocked(t, "fakeUUID1", dIn)
	GetDispatcherScriptCharacterizationsByUUIDFailErrMocked(t, "fakeUUID1", dIn)
	GetDispatcherScriptCharacterizationsByUUIDFailStatusMocked(t, "fakeUUID1", dIn)
	GetDispatcherScriptCharacterizationsByUUIDFailJSONMocked(t, "fakeUUID1", dIn)
}

func TestReportScriptConclusions(t *testing.T) {
	dIn := testdata.GetScriptConclusionData()
	ReportScriptConclusionsMocked(t, dIn)
	ReportScriptConclusionsFailErrMocked(t, dIn)
	ReportScriptConclusionsFailStatusMocked(t, dIn)
	ReportScriptConclusionsFailJSONMocked(t, dIn)
}

func TestDownloadAttachment(t *testing.T) {
	dataIn := testdata.GetDownloadAttachmentData()
	DownloadAttachmentMocked(t, dataIn)
	DownloadAttachmentFailErrMocked(t, dataIn)
}
