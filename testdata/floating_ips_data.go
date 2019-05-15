package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetFloatingIPData loads test data
func GetFloatingIPData() []*types.FloatingIP {
	return []*types.FloatingIP{
		{
			ID:               "fakeId0",
			Name:             "fakeName0",
			Address:          "fakeAddress0",
			State:            "fakeState0",
			CloudAccountID:   "fakeCloudAccountID0",
			RealmID:          "fakeRealmID0",
			AttachedServerID: "fakeAttachedServerID0",
			Brownfield:       false,
		},
		{
			ID:               "fakeId1",
			Name:             "fakeName1",
			Address:          "fakeAddress1",
			State:            "fakeState1",
			CloudAccountID:   "fakeCloudAccountID1",
			RealmID:          "fakeRealmID1",
			AttachedServerID: "fakeAttachedServerID1",
			Brownfield:       false,
		},
	}
}
