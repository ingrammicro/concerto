package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetStoragePlanData loads test data
func GetStoragePlanData() []*types.StoragePlan {

	return []*types.StoragePlan{
		{
			ID:                  "fakeID0",
			Name:                "fakeName0",
			MinSize:             1,
			MaxSize:             10,
			CloudProviderID:     "fakeCloudProviderID0",
			CloudProviderName:   "fakeCloudProviderName0",
			LocationID:          "fakeLocationID0",
			LocationName:        "fakeLocationName0",
			RealmID:             "fakeRealmID0",
			RealmProviderName:   "fakeRealmProviderName0",
			FlavourProviderName: "fakeFlavourProviderName0",
			Deprecated:          false,
		},
		{
			ID:                  "fakeID1",
			Name:                "fakeName1",
			MinSize:             1,
			MaxSize:             65536,
			CloudProviderID:     "fakeCloudProviderID1",
			CloudProviderName:   "fakeCloudProviderName1",
			LocationID:          "fakeLocationID1",
			LocationName:        "fakeLocationName1",
			RealmID:             "fakeRealmID1",
			RealmProviderName:   "fakeRealmProviderName1",
			FlavourProviderName: "fakeFlavourProviderName1",
			Deprecated:          true,
		},
	}
}
