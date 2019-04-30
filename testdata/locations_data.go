package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetLocationData loads test data
func GetLocationData() []*types.Location {

	return []*types.Location{
		{
			ID:   "fakeID0",
			Name: "fakeName0",
		},
		{
			ID:   "fakeID1",
			Name: "fakeName1",
		},
	}
}
