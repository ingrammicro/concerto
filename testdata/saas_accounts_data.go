package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetSaasAccountData loads test data
func GetSaasAccountData() *[]types.SaasAccount {

	testSaasAccounts := []types.SaasAccount{
		{
			ID:             "fakeID0",
			SaasProviderID: "fakeSaasProviderID0",
		},
		{
			ID:             "fakeID1",
			SaasProviderID: "fakeSaasProviderID1",
		},
	}

	return &testSaasAccounts
}
