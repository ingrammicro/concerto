package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/network"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpFirewallProfile prepares common resources to send request to Concerto API
func WireUpFirewallProfile(c *cli.Context) (ds *network.FirewallProfileService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewFirewallProfileService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up firewallProfile service", err)
	}

	return ds, f
}

// FirewallProfileList subcommand function
func FirewallProfileList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	firewallProfiles, err := firewallProfileSvc.GetFirewallProfileList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewallProfile data", err)
	}

	labelables := make([]types.Labelable, len(firewallProfiles))
	for i, fwp := range firewallProfiles {
		labelables[i] = types.Labelable(&fwp)
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)
	firewallProfiles = make([]types.FirewallProfile, len(filteredLabelables))
	for i, labelable := range labelables {
		fw, ok := labelable.(*types.FirewallProfile)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.FirewallProfile, got a %T", labelable))
		}
		firewallProfiles[i] = *fw
	}
	if err = formatter.PrintList(firewallProfiles); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FirewallProfileShow subcommand function
func FirewallProfileShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	firewallProfile, err := firewallProfileSvc.GetFirewallProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewallProfile data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	firewallProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FirewallProfileCreate subcommand function
func FirewallProfileCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"name", "description"}, formatter)
	params, err := utils.FlagConvertParamsJSON(c, []string{"rules"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	firewallProfileIn := map[string]interface{}{
		"name":        c.String("name"),
		"description": c.String("description"),
	}
	if c.String("rules") != "" {
		firewallProfileIn["rules"] = (*params)["rules"]
	}
	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"))
		firewallProfileIn["label_ids"] = labelsIdsArr
	}

	firewallProfile, err := firewallProfileSvc.CreateFirewallProfile(&firewallProfileIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create firewallProfile", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	firewallProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FirewallProfileUpdate subcommand function
func FirewallProfileUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	params, err := utils.FlagConvertParamsJSON(c, []string{"rules"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}
	firewallProfile, err := firewallProfileSvc.UpdateFirewallProfile(params, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update firewallProfile", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	firewallProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FirewallProfileDelete subcommand function
func FirewallProfileDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := firewallProfileSvc.DeleteFirewallProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete firewallProfile", err)
	}
	return nil
}
