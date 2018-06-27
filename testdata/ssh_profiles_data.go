package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetSSHProfileData loads test data
func GetSSHProfileData() *[]types.SSHProfile {

	testSSHProfiles := []types.SSHProfile{
		{
			ID:         "fakeID0",
			Name:       "fakeName0",
			PublicKey:  "fakePublicKey0",
			PrivateKey: "fakePrivateKey0",
		},
		{
			ID:         "fakeID1",
			Name:       "fakeName1",
			PublicKey:  "fakePublicKey1",
			PrivateKey: "fakePrivateKey1",
		},
	}

	return &testSSHProfiles
}
