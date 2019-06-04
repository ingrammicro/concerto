package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetServerArrayData loads test data
func GetServerArrayData() []*types.ServerArray {

	return []*types.ServerArray{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			State:             "fakeState0",
			Size:              0,
			TemplateID:        "fakeTemplateID0",
			CloudAccountID:    "fakeCloudAccountID0",
			ServerPlanID:      "fakeServerPlanID0",
			FirewallProfileID: "fakeFirewallProfileID0",
			SSHProfileID:      "fakeSSHProfileID0",
			SubnetID:          "fakeSubnetID0",
			VpcID:             "fakeVpcID0",
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			State:             "fakeState1",
			Size:              2,
			TemplateID:        "fakeTemplateID1",
			CloudAccountID:    "fakeCloudAccountID1",
			ServerPlanID:      "fakeServerPlanID1",
			FirewallProfileID: "fakeFirewallProfileID1",
			SSHProfileID:      "fakeSSHProfileID1",
			SubnetID:          "fakeSubnetID1",
			VpcID:             "fakeVpcID1",
			Privateness:       true,
		},
	}
}
