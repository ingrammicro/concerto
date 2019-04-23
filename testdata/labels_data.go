package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetLabelData loads test data
func GetLabelData() []*types.Label {

	return []*types.Label{
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
}

// GetLabelWithNamespaceData loads test data
func GetLabelWithNamespaceData() []*types.Label {

	return []*types.Label{
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
}

// GetLabeledResourcesData loads test data
func GetLabeledResourcesData() []*types.LabeledResource {

	return []*types.LabeledResource{
		{
			ID:           "fakeID0",
			ResourceType: "server",
		},
		{
			ID:           "fakeID1",
			ResourceType: "template",
		},
	}
}
