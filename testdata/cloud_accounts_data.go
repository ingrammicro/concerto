package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetCloudAccountData loads test data
func GetCloudAccountData() *[]types.CloudAccount {

	testCloudAccounts := []types.CloudAccount{
		{
			Id:            "fakeID0",
			Name:          "fakeName0",
			CloudProvId:   "fakeProvID0",
			CloudProvName: "fakeProvName0",
		},
		{
			Id:            "fakeID1",
			Name:          "fakeName1",
			CloudProvId:   "fakeProvID1",
			CloudProvName: "fakeProvName1",
		},
	}

	return &testCloudAccounts
}
