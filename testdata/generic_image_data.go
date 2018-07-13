package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetGenericImageData loads test data
func GetGenericImageData() *[]types.GenericImage {

	testGenericImages := []types.GenericImage{
		{
			ID:   "fakeID0",
			Name: "fakeName0",
		},
		{
			ID:   "fakeID1",
			Name: "fakeName1",
		},
	}

	return &testGenericImages
}
