package labels

import (
	"testing"

	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewLabelServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewLabelService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCloudProviderList(t *testing.T) {
	labelsIn := testdata.GetLabelData()
	GetLabelListMocked(t, labelsIn)
	GetLabelListFailErrMocked(t, labelsIn)
	GetLabelListFailStatusMocked(t, labelsIn)
	GetLabelListFailJSONMocked(t, labelsIn)
}

func TestCreateLabel(t *testing.T) {
	labelsIn := testdata.GetLabelData()
	for _, labelIn := range *labelsIn {
		CreateLabelMocked(t, &labelIn)
		CreateLabelFailErrMocked(t, &labelIn)
		CreateLabelFailStatusMocked(t, &labelIn)
		//CreateLabelFailJSONMocked(t, &labelIn)
	}
}
