package testdata

import (
	"encoding/json"

	"github.com/ingrammicro/concerto/api/types"
)

// GetTemplateData loads loads test data
func GetTemplateData() []*types.Template {
	conf0 := json.RawMessage(`{"fakeConf01":"x","fakeConf02":"y"}`)
	conf1 := json.RawMessage(`{"fakeConf11":"x","fakeConf12":"y"}`)

	return []*types.Template{
		{
			ID:                      "fakeID0",
			Name:                    "fakeName0",
			GenericImageID:          "fakeGenericImageID0",
			RunList:                 []string{"fakeRunList01", "fakeRunList02"},
			ConfigurationAttributes: &conf0,
		},
		{
			ID:                      "fakeID1",
			Name:                    "fakeName1",
			GenericImageID:          "fakeGenericImageID1",
			RunList:                 []string{"fakeRunList11", "fakeRunList12", "fakeRunList13"},
			ConfigurationAttributes: &conf1,
		},
		{
			ID:                      "fakeID2",
			Name:                    "fakeName2",
			GenericImageID:          "fakeGenericImageID2",
			RunList:                 []string{"fakeRunList21", "fakeRunList22", "fakeRunList23"},
			ConfigurationAttributes: nil,
		},
	}
}

// GetTemplateScriptData loads test data
func GetTemplateScriptData() []*types.TemplateScript {

	param0 := json.RawMessage(`{"fakeConf01":"x","fakeConf02":"y"}`)
	param1 := json.RawMessage(`{"fakeConf11":"x","fakeConf12":"y"}`)
	param2 := json.RawMessage(`{"fakeConf21":"x","fakeConf22":"y"}`)
	param3 := json.RawMessage(`{"fakeConf31":"x","fakeConf32":"y"}`)

	return []*types.TemplateScript{
		{
			ID:              "fakeID0",
			Type:            "fakeType0",
			ExecutionOrder:  1,
			TemplateID:      "fakeTemplateID0",
			ScriptID:        "fakeScriptID0",
			ParameterValues: &param0,
		},
		{
			ID:              "fakeID1",
			Type:            "fakeType1",
			ExecutionOrder:  4,
			TemplateID:      "fakeTemplateID1",
			ScriptID:        "fakeScriptID1",
			ParameterValues: &param1,
		},
		{
			ID:              "fakeID2",
			Type:            "fakeType2",
			ExecutionOrder:  2,
			TemplateID:      "fakeTemplateID2",
			ScriptID:        "fakeScriptID2",
			ParameterValues: &param2,
		},
		{
			ID:              "fakeID3",
			Type:            "fakeType3",
			ExecutionOrder:  3,
			TemplateID:      "fakeTemplateID3",
			ScriptID:        "fakeScriptID3",
			ParameterValues: &param3,
		},
	}
}

// GetTemplateServerData loads loads test data
func GetTemplateServerData() []*types.TemplateServer {

	return []*types.TemplateServer{
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
