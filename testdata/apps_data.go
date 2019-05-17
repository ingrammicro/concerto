package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetAppData loads test data
func GetAppData() []*types.WizardApp {
	return []*types.WizardApp{
		{
			ID:                  "fakeID0",
			Name:                "fakeName0",
			FlavourRequirements: map[string]interface{}{"fakeFlavour01": "x", "fakeFlavour02": "y"},
			GenericImageID:      "fakeGenericImageID0",
		},
		{
			ID:                  "fakeID1",
			Name:                "fakeName1",
			FlavourRequirements: map[string]interface{}{"fakeFlavour11": "a", "fakeFlavour12": "b"},
			GenericImageID:      "fakeGenericImageID1",
		},
	}
}
