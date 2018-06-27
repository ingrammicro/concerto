package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetLocationData loads test data
func GetLocationData() *[]types.Location {

	testLocations := []types.Location{
		{
			ID:   "fakeID0",
			Name: "fakeName0",
		},
		{
			ID:   "fakeID1",
			Name: "fakeName1",
		},
	}

	return &testLocations
}
