package testdata

import (
	"encoding/json"
	"github.com/ingrammicro/concerto/api/types"
)

// GetServerData loads test data
func GetServerData() []*types.Server {

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
}

// GetScriptCharData loads test data
func GetScriptCharData() []*types.ScriptChar {

	param0 := json.RawMessage(`{"fakeConf01":"x","fakeConf02":"y"}`)
	param1 := json.RawMessage(`{"fakeConf11":"x","fakeConf12":"y"}`)

	return []*types.ScriptChar{
		{
			ID:              "fakeID0",
			Type:            "fakeType0",
			ParameterValues: &param0,
			TemplateID:      "fakeTemplateID0",
			ScriptID:        "fakeScriptID0",
			ExecutionOrder:  0,
			ResourceType:    "fakeResourceType0",
		},
		{
			ID:              "fakeID1",
			Type:            "fakeType1",
			ParameterValues: &param1,
			TemplateID:      "fakeTemplateID1",
			ScriptID:        "fakeScriptID1",
			ExecutionOrder:  1,
			ResourceType:    "fakeResourceType1",
		},
	}
}
