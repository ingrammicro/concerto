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

// GetSubnetServersData loads test data
func GetSubnetServersData() []*types.Server {
	return []*types.Server{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			Fqdn:              "fakeFqdn0",
			State:             "fakeState0",
			PublicIP:          "fakePublicIP0",
			TemplateID:        "fakeTemplateID0",
			ServerPlanID:      "fakeServerPlanID0",
			SSHProfileID:      "fakeSSHProfileID0",
			FirewallProfileID: "fakeFirewallProfileID0",
			SubnetID:          "fakeSubnetID0",
			VpcID:             "fakeVpcID0",
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			Fqdn:              "fakeFqdn1",
			State:             "fakeState1",
			PublicIP:          "fakePublicIP1",
			TemplateID:        "fakeTemplateID1",
			ServerPlanID:      "fakeServerPlanID1",
			SSHProfileID:      "fakeSSHProfileID1",
			FirewallProfileID: "fakeFirewallProfileID1",
			SubnetID:          "fakeSubnetID1",
			VpcID:             "fakeVpcID1",
		},
	}
}
