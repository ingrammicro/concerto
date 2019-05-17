package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetSubnetData loads test data
func GetSubnetData() []*types.Subnet {
	return []*types.Subnet{
		{
			ID:                     "fakeId0",
			Name:                   "fakeName0",
			CIDR:                   "fakeCIDR0",
			State:                  "fakeState0",
			Type:                   "fakeType0",
			VpcID:                  "fakeVpcID0",
			ServerCreationDisabled: false,
			Brownfield:             false,
		},
		{
			ID:                     "fakeId1",
			Name:                   "fakeName1",
			CIDR:                   "fakeCIDR1",
			State:                  "fakeState1",
			Type:                   "fakeType1",
			VpcID:                  "fakeVpcID1",
			ServerCreationDisabled: false,
			Brownfield:             false,
		},
	}
}
