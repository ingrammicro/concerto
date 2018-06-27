package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetServerData loads test data
func GetServerData() *[]types.Server {

	testServers := []types.Server{
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
		},
	}

	return &testServers
}

// GetDNSData loads test data
func GetDNSData() *[]types.Dns {

	testDnss := []types.Dns{
		{
			ID:       "fakeID0",
			Name:     "fakeName0",
			Content:  "fakeContent0",
			Type:     "fakeType0",
			IsFQDN:   true,
			DomainID: "fakeDomainID0",
		},
		{
			ID:       "fakeID1",
			Name:     "fakeName1",
			Content:  "fakeContent1",
			Type:     "fakeType1",
			IsFQDN:   false,
			DomainID: "fakeDomainID1",
		},
	}

	return &testDnss
}

// GetScriptCharData loads test data
func GetScriptCharData() *[]types.ScriptChar {

	testScriptChars := []types.ScriptChar{
		{
			ID:   "fakeID0",
			Type: "fakeType0",
			// Parameter_values: struct{"fakeInst0", "fakeInst1"},
			TemplateID: "fakeTemplateID0",
			ScriptID:   "fakeScriptID0",
		},
		{
			ID:   "fakeID1",
			Type: "fakeType1",
			// Parameter_values: struct{"fakeInst2", "fakeInst2", "fakeInst3"},
			TemplateID: "fakeTemplateID1",
			ScriptID:   "fakeScriptID1",
		},
	}

	return &testScriptChars
}
