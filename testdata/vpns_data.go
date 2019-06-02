package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetVPNData loads test data
func GetVPNData() []*types.Vpn {
	return []*types.Vpn{
		{
			ID:           "fakeId0",
			State:        "fakeState0",
			VpcID:        "fakeVpcID0",
			VpnPlanID:    "fakeVpcPlanID0",
			PublicIP:     "fakePublicIP0",
			ExposedCIDRs: []string{"fakeCIDR00", "fakeCIDR01"},
		},
		{
			ID:           "fakeId1",
			State:        "fakeState1",
			VpcID:        "fakeVpcID1",
			VpnPlanID:    "fakeVpcPlanID1",
			PublicIP:     "fakePublicIP1",
			ExposedCIDRs: []string{"fakeCIDR10", "fakeCIDR11"},
		},
	}
}

// GetVPNPlanData loads test data
func GetVPNPlanData() []*types.VpnPlan {
	return []*types.VpnPlan{
		{
			ID:       "fakeId0",
			Name:     "fakeName0",
			Active:   "fakeActive0",
			RemoteID: "fakeRemoteID0",
		},
		{
			ID:       "fakeId1",
			Name:     "fakeName1",
			Active:   "fakeActive1",
			RemoteID: "fakeRemoteID1",
		},
	}
}
