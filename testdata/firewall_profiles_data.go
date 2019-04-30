package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetFirewallProfileData loads test data
func GetFirewallProfileData() []*types.FirewallProfile {

	return []*types.FirewallProfile{
		{
			ID:          "fakeId0",
			Name:        "fakeName0",
			Description: "fakeDescription0",
			Default:     true,
			Rules: []types.Rule{
				{
					Protocol: "fakeProtocol0",
					MinPort:  0,
					MaxPort:  1024,
					CidrIP:   "fakeCidrIP0",
				},
			},
		},
		{
			ID:          "fakeId1",
			Name:        "fakeName1",
			Description: "fakeDescription1",
			Default:     false,
			Rules: []types.Rule{
				{
					Protocol: "fakeProtocol1",
					MinPort:  0,
					MaxPort:  200,
					CidrIP:   "fakeCidrIP1",
				},
				{
					Protocol: "fakeProtocol2",
					MinPort:  0,
					MaxPort:  2048,
					CidrIP:   "fakeCidrIP2",
				},
			},
		},
	}
}
