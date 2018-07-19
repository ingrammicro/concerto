package labels

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// GetLabelListMocked test mocked function
func GetLabelListMocked(t *testing.T, labelsIn *[]types.Label) *[]types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelsIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Get", "/v1/labels").Return(dIn, 200, nil)
	labelsOut, err := ds.GetLabelList()
	assert.Nil(err, "Error getting labels list")
	assert.Equal(*labelsIn, labelsOut, "GetLabelList returned different labels")

	return &labelsOut
}

// GetLabelListMockedWithNamespace test mocked function
func GetLabelListMockedWithNamespace(t *testing.T, labelsIn *[]types.Label) *[]types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelsIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Get", "/v1/labels").Return(dIn, 200, nil)
	labelsOut, err := ds.GetLabelList()
	assert.Nil(err, "Error getting labels list")
	assert.NotEqual(*labelsIn, labelsOut, "GetLabelList returned labels with Namespaces")

	return &labelsOut
}

// GetLabelListFailErrMocked test mocked function
func GetLabelListFailErrMocked(t *testing.T, labelsIn *[]types.Label) *[]types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelsIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Get", "/v1/labels").Return(dIn, 200, fmt.Errorf("Mocked error"))
	labelsOut, err := ds.GetLabelList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(labelsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &labelsOut
}

// GetLabelListFailStatusMocked test mocked function
func GetLabelListFailStatusMocked(t *testing.T, labelsIn *[]types.Label) *[]types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelsIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Get", "/v1/labels").Return(dIn, 499, nil)
	labelsOut, err := ds.GetLabelList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(labelsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &labelsOut
}

// GetLabelListFailJSONMocked test mocked function
func GetLabelListFailJSONMocked(t *testing.T, labelsIn *[]types.Label) *[]types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/labels").Return(dIn, 200, nil)
	labelsOut, err := ds.GetLabelList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(labelsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &labelsOut
}

// CreateLabelMocked test mocked function
func CreateLabelMocked(t *testing.T, labelIn *types.Label) *types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// to json
	dOut, err := json.Marshal(labelIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Post", "/v1/labels/", mapIn).Return(dOut, 200, nil)
	labelOut, err := ds.CreateLabel(mapIn)
	assert.Nil(err, "Error creating label")
	assert.Equal(labelIn, labelOut, "CreateLabel returned different labels")

	return labelOut
}

// CreateLabelFailErrMocked test mocked function
func CreateLabelFailErrMocked(t *testing.T, labelIn *types.Label) *types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// to json
	dOut, err := json.Marshal(labelIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Post", "/v1/labels/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	labelOut, err := ds.CreateLabel(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(labelOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return labelOut
}

// CreateLabelFailStatusMocked test mocked function
func CreateLabelFailStatusMocked(t *testing.T, labelIn *types.Label) *types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// to json
	dOut, err := json.Marshal(labelIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Post", "/v1/labels/", mapIn).Return(dOut, 499, nil)
	labelOut, err := ds.CreateLabel(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(labelOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return labelOut
}

// CreateLabelFailJSONMocked test mocked function
func CreateLabelFailJSONMocked(t *testing.T, labelIn *types.Label) *types.Label {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/labels/", mapIn).Return(dIn, 200, nil)
	labelOut, err := ds.CreateLabel(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(labelOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return labelOut
}

// AddLabelMocked test mocked function
func AddLabelMocked(t *testing.T, labelIn *types.Label, labeledResourcesOut []types.LabeledResources) []types.LabeledResources {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// to json
	dOut, err := json.Marshal(labeledResourcesOut)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v1/labels/%s/resources", labelIn.ID), mapIn).Return(dOut, 200, nil)
	labeledOut, err := ds.AddLabel(mapIn, labelIn.ID)
	assert.Nil(err, "Error creating label")
	assert.Equal(labeledOut, labeledResourcesOut, "CreateLabel returned invalid labeled resources")

	return labeledResourcesOut
}

// AddLabelFailErrMocked test mocked function
func AddLabelFailErrMocked(t *testing.T, labelIn *types.Label, labeledResourcesOut []types.LabeledResources) []types.LabeledResources {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// to json
	dOut, err := json.Marshal(labeledResourcesOut)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v1/labels/%s/resources", labelIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	labeledOut, err := ds.AddLabel(mapIn, labelIn.ID)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(labeledOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return labeledResourcesOut
}

// AddLabelFailStatusMocked test mocked function
func AddLabelFailStatusMocked(t *testing.T, labelIn *types.Label, labeledResourcesOut []types.LabeledResources) []types.LabeledResources {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// to json
	dOut, err := json.Marshal(labeledResourcesOut)
	assert.Nil(err, "Label test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v1/labels/%s/resources", labelIn.ID), mapIn).Return(dOut, 404, nil)
	labeledOut, err := ds.AddLabel(mapIn, labelIn.ID)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(labeledOut, "Expecting nil output")
	assert.Contains(err.Error(), "404", "Error should contain http code 404")

	return labeledResourcesOut
}

// AddLabelFailJSONMocked test mocked function
func AddLabelFailJSONMocked(t *testing.T, labelIn *types.Label, labeledResourcesOut []types.LabeledResources) []types.LabeledResources {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*labelIn)
	assert.Nil(err, "Label test data corrupted")

	// wrong json
	dOut := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/v1/labels/%s/resources", labelIn.ID), mapIn).Return(dOut, 200, nil)
	labeledOut, err := ds.AddLabel(mapIn, labelIn.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(labeledOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return labeledResourcesOut
}

// RemoveLabelMocked test mocked function
func RemoveLabelMocked(t *testing.T, labelIn *types.Label) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	resourceID := "5b5074735f7c880ad9c6bbce"
	cs.On("Delete", fmt.Sprintf("v1/labels/%s/resources/%s/%s", labelIn.ID, labelIn.ResourceType, resourceID)).Return(dIn, 204, nil)
	err = ds.RemoveLabel(labelIn.ID, labelIn.ResourceType, resourceID)
	assert.Nil(err, "Error removing label")
}

// RemoveLabelFailErrMocked test mocked function
func RemoveLabelFailErrMocked(t *testing.T, labelIn *types.Label) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	resourceID := "5b5074735f7c880ad9c6bbce"
	cs.On("Delete", fmt.Sprintf("v1/labels/%s/resources/%s/%s", labelIn.ID, labelIn.ResourceType, resourceID)).Return(dIn, 204, fmt.Errorf("Mocked error"))
	err = ds.RemoveLabel(labelIn.ID, labelIn.ResourceType, resourceID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// RemoveLabelFailStatusMocked test mocked function
func RemoveLabelFailStatusMocked(t *testing.T, labelIn *types.Label) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLabelService(cs)
	assert.Nil(err, "Couldn't load label service")
	assert.NotNil(ds, "Label service not instanced")

	// to json
	dIn, err := json.Marshal(labelIn)
	assert.Nil(err, "Label test data corrupted")

	// call service
	resourceID := "5b5074735f7c880ad9c6bbce"
	cs.On("Delete", fmt.Sprintf("v1/labels/%s/resources/%s/%s", labelIn.ID, labelIn.ResourceType, resourceID)).Return(dIn, 404, nil)
	err = ds.RemoveLabel(labelIn.ID, labelIn.ResourceType, resourceID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "404", "Error should contain http code 404")
}
