package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetCloudProviderData loads test data
func GetCloudProviderData() []*types.CloudProvider {

	return []*types.CloudProvider{
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
