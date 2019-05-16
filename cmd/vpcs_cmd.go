package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/network"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpVPC prepares common resources to send request to Concerto API
func WireUpVPC(c *cli.Context) (ds *network.VPCService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewVPCService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up VPC service", err)
	}

	return ds, f
}

// VPCList subcommand function
func VPCList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpcSvc, formatter := WireUpVPC(c)

	vpcs, err := vpcSvc.GetVPCList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive VPC data", err)
	}

	labelables := make([]types.Labelable, len(vpcs))
	for i := 0; i < len(vpcs); i++ {
		labelables[i] = types.Labelable(vpcs[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	vpcs = make([]*types.Vpc, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Vpc)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Vpc, got a %T", labelable))
		}
		vpcs[i] = v
	}
	if err = formatter.PrintList(vpcs); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VPCShow subcommand function
func VPCShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpcSvc, formatter := WireUpVPC(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	vpc, err := vpcSvc.GetVPC(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive VPC data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VPCCreate subcommand function
func VPCCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpcSvc, formatter := WireUpVPC(c)

	checkRequiredFlags(c, []string{"name", "cidr", "cloud-account-id", "realm-provider-name"}, formatter)

	vpcIn := map[string]interface{}{
		"name":                c.String("name"),
		"cidr":                c.String("cidr"),
		"cloud_account_id":    c.String("cloud-account-id"),
		"realm_provider_name": c.String("realm-provider-name"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		vpcIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	vpc, err := vpcSvc.CreateVPC(&vpcIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create VPC", err)
	}

	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VPCUpdate subcommand function
func VPCUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpcSvc, formatter := WireUpVPC(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)

	vpcIn := map[string]interface{}{
		"name": c.String("name"),
	}

	vpc, err := vpcSvc.UpdateVPC(&vpcIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update VPC", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VPCDelete subcommand function
func VPCDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpcSvc, formatter := WireUpVPC(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := vpcSvc.DeleteVPC(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete VPC", err)
	}
	return nil
}
