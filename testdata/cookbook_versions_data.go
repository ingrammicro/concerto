package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetCookbookVersionData loads test data
func GetCookbookVersionData() *[]types.CookbookVersion {
	testCookbookVersions := []types.CookbookVersion{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			Description:       "fakeDescription0",
			Version:           "fakeVersion0",
			State:             "fakeState0",
			RevisionID:        "fakeRevisionID0",
			Recipes:           []string{"fakeRecipe01", "fakeRecipe02"},
			ResourceType:      "fakeResourceType0",
			PubliclyAvailable: true,
			GlobalLegacy:      true,
			UploadURL:         "fakeUploadURL0",
			ErrorMessage:      "",
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			Description:       "fakeDescription1",
			Version:           "fakeVersion1",
			State:             "fakeState1",
			RevisionID:        "fakeRevisionID1",
			Recipes:           []string{"fakeRecipe11", "fakeRecipe12"},
			ResourceType:      "fakeResourceType1",
			PubliclyAvailable: true,
			GlobalLegacy:      true,
			UploadURL:         "fakeUploadURL1",
			ErrorMessage:      "",
		},
	}

	return &testCookbookVersions
}
