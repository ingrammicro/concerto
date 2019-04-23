package blueprint

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCookbookVersionServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCookbookVersionService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCookbookVersionList(t *testing.T) {
	cbsIn := testdata.GetCookbookVersionData()
	GetCookbookVersionListMocked(t, cbsIn)
	GetCookbookVersionListFailErrMocked(t, cbsIn)
	GetCookbookVersionListFailStatusMocked(t, cbsIn)
	GetCookbookVersionListFailJSONMocked(t, cbsIn)
}

func TestGetCookbookVersion(t *testing.T) {
	cbsIn := testdata.GetCookbookVersionData()
	for _, cbIn := range cbsIn {
		GetCookbookVersionMocked(t, cbIn)
		GetCookbookVersionFailErrMocked(t, cbIn)
		GetCookbookVersionFailStatusMocked(t, cbIn)
		GetCookbookVersionFailJSONMocked(t, cbIn)
	}
}

func TestCreateCookbookVersion(t *testing.T) {
	cbsIn := testdata.GetCookbookVersionData()
	for _, cbIn := range cbsIn {
		CreateCookbookVersionMocked(t, cbIn)
		CreateCookbookVersionFailErrMocked(t, cbIn)
		CreateCookbookVersionFailStatusMocked(t, cbIn)
		CreateCookbookVersionFailJSONMocked(t, cbIn)
	}
}

func TestUploadCookbookVersion(t *testing.T) {
	cbsIn := testdata.GetCookbookVersionData()
	for _, cbIn := range cbsIn {
		UploadCookbookVersionMocked(t, cbIn)
		UploadCookbookVersionFailStatusMocked(t, cbIn)
		UploadCookbookVersionFailErrMocked(t, cbIn)
	}
}

func TestProcessCookbookVersion(t *testing.T) {
	cbsIn := testdata.GetCookbookVersionData()
	for _, cbIn := range cbsIn {
		ProcessCookbookVersionMocked(t, cbIn)
		ProcessCookbookVersionFailErrMocked(t, cbIn)
		ProcessCookbookVersionFailStatusMocked(t, cbIn)
		ProcessCookbookVersionFailJSONMocked(t, cbIn)
	}
}

func TestDeleteCookbookVersion(t *testing.T) {
	cbsIn := testdata.GetCookbookVersionData()
	for _, cbIn := range cbsIn {
		DeleteCookbookVersionMocked(t, cbIn)
		DeleteCookbookVersionFailErrMocked(t, cbIn)
		DeleteCookbookVersionFailStatusMocked(t, cbIn)
	}
}
