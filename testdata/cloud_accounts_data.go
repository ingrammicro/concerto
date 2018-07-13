package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetCloudAccountData loads test data
func GetCloudAccountData() *[]types.CloudAccount {

	testCloudAccounts := []types.CloudAccount{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			CloudProviderID:   "CloudProviderID0",
			CloudProviderName: "CloudProviderName0",
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			CloudProviderID:   "CloudProviderID1",
			CloudProviderName: "CloudProviderName1",
		},
	}

	return &testCloudAccounts
}
