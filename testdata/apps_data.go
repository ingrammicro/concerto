package testdata

import (
	"encoding/json"

	"github.com/ingrammicro/concerto/api/types"
)

// GetAppData loads test data
func GetAppData() *[]types.WizardApp {

	param0 := json.RawMessage(`{"fakeFlavour01":"x","fakeFlavour02":"y"}`)
	param1 := json.RawMessage(`{"fakeFlavour11":"a","fakeFlavour12":"b"}`)

	testApps := []types.WizardApp{
		{
			ID:                  "fakeID0",
			Name:                "fakeName0",
			FlavourRequirements: param0,
			GenericImageID:      "fakeGenericImageID0",
		},
		{
			ID:                  "fakeID1",
			Name:                "fakeName1",
			FlavourRequirements: param1,
			GenericImageID:      "fakeGenericImageID1",
		},
	}

	return &testApps
}
