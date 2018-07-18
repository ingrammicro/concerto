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
