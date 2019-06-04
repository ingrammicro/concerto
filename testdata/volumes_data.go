package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetVolumeData loads test data
func GetVolumeData() []*types.Volume {

	return []*types.Volume{
		{
			ID:               "fakeID0",
			Name:             "fakeName0",
			Size:             1,
			State:            "fakeState0",
			Device:           "fakeDevice0",
			StoragePlanID:    "fakeStoragePlanID0",
			CloudAccountID:   "fakeCloudAccountID0",
			RealmID:          "fakeRealmID0",
			AttachedServerID: "fakeAttachedServerID0",
			Brownfield:       false,
		},
		{
			ID:               "fakeID1",
			Name:             "fakeName1",
			Size:             2,
			State:            "fakeState1",
			Device:           "fakeDevice1",
			StoragePlanID:    "fakeStoragePlanID1",
			CloudAccountID:   "fakeCloudAccountID1",
			RealmID:          "fakeRealmID1",
			AttachedServerID: "fakeAttachedServerID1",
			Brownfield:       false,
		},
	}
}
