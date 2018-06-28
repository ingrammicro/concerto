package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetSaasProviderData loads test data
func GetSaasProviderData() *[]types.SaasProvider {

	testSaasProviders := []types.SaasProvider{
		{
			ID:                  "fakeID0",
			Name:                "fakeName0",
			RequiredAccountData: []string{"accData0"},
		},
		{
			ID:                  "fakeID1",
			Name:                "fakeName1",
			RequiredAccountData: []string{"accData1", "accData2", "accData3"},
		},
	}

	return &testSaasProviders
}
