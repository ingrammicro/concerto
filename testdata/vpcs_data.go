package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetVPCData loads test data
func GetVPCData() []*types.Vpc {
	return []*types.Vpc{
		{
			ID:                 "fakeId0",
			Name:               "fakeName0",
			CIDR:               "fakeCIDR0",
			State:              "fakeState0",
			CloudAccountID:     "fakeCloudAccountID0",
			RealmProviderName:  "fakeRealmProviderName0",
			HasVPN:             false,
			AllowedSubnetTypes: []string{"fakeSubNet00", "fakeSubNet01"},
			Brownfield:         false,
		},
		{
			ID:                 "fakeId1",
			Name:               "fakeName1",
			CIDR:               "fakeCIDR1",
			State:              "fakeState1",
			CloudAccountID:     "fakeCloudAccountID1",
			RealmProviderName:  "fakeRealmProviderName1",
			HasVPN:             false,
			AllowedSubnetTypes: []string{"fakeSubNet10", "fakeSubNet11"},
			Brownfield:         false,
		},
	}
}
