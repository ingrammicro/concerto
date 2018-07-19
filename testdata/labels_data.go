package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetLabelData loads test data
func GetLabelData() *[]types.Label {

	testLabels := []types.Label{
		{
			ID:           "fakeID0",
			Name:         "fakeName0",
			ResourceType: "label",
		},
		{
			ID:           "fakeID1",
			Name:         "fakeName1",
			ResourceType: "label",
		},
	}

	return &testLabels
}

// GetLabelWithNamespaceData loads test data
func GetLabelWithNamespaceData() *[]types.Label {

	testLabels := []types.Label{
		{
			ID:           "fakeID0",
			Name:         "fakeName0",
			ResourceType: "label",
			Namespace:    "fakeNamespace0",
			Value:        "fakeValue0",
		},
		{
			ID:           "fakeID1",
			Name:         "fakeName1",
			ResourceType: "label",
			Namespace:    "fakeNamespace1",
			Value:        "fakeValue1",
		},
	}

	return &testLabels
}

// GetLabeledResourcesData loads test data
func GetLabeledResourcesData() *[]types.LabeledResources {

	testLabeledResources := []types.LabeledResources{
		{
			ID:           "fakeID0",
			ResourceType: "server",
		},
		{
			ID:           "fakeID1",
			ResourceType: "template",
		},
	}

	return &testLabeledResources
}
