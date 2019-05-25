package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/network"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpFloatingIP prepares common resources to send request to Concerto API
func WireUpFloatingIP(c *cli.Context) (ds *network.FloatingIPService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewFloatingIPService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up floating IP service", err)
	}

	return ds, f
}

// FloatingIPList subcommand function
func FloatingIPList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	floatingIPs, err := floatingIPSvc.GetFloatingIPList(c.String("server-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive floating IP data", err)
	}

	labelables := make([]types.Labelable, len(floatingIPs))
	for i := 0; i < len(floatingIPs); i++ {
		labelables[i] = types.Labelable(floatingIPs[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	floatingIPs = make([]*types.FloatingIP, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		fw, ok := labelable.(*types.FloatingIP)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.FloatingIP, got a %T", labelable))
		}
		floatingIPs[i] = fw
	}
	if err = formatter.PrintList(floatingIPs); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FloatingIPShow subcommand function
func FloatingIPShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	floatingIP, err := floatingIPSvc.GetFloatingIP(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive floating IP data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FloatingIPCreate subcommand function
func FloatingIPCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"name", "cloud-account-id", "realm-id"}, formatter)

	floatingIPIn := map[string]interface{}{
		"name":             c.String("name"),
		"cloud_account_id": c.String("cloud-account-id"),
		"realm_id":         c.String("realm-id"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		floatingIPIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	floatingIP, err := floatingIPSvc.CreateFloatingIP(&floatingIPIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create floating IP", err)
	}

	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FloatingIPUpdate subcommand function
func FloatingIPUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)

	floatingIPIn := map[string]interface{}{
		"name": c.String("name"),
	}

	floatingIP, err := floatingIPSvc.UpdateFloatingIP(&floatingIPIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update floating IP", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FloatingIPAttach subcommand function
func FloatingIPAttach(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"id", "server-id"}, formatter)

	floatingIPIn := map[string]interface{}{
		"attached_server_id": c.String("server-id"),
	}

	server, err := floatingIPSvc.AttachFloatingIP(&floatingIPIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't attach floating IP", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FloatingIPDetach subcommand function
func FloatingIPDetach(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := floatingIPSvc.DetachFloatingIP(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't detach floating IP", err)
	}
	return nil
}

// FloatingIPDelete subcommand function
func FloatingIPDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := floatingIPSvc.DeleteFloatingIP(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete floating IP", err)
	}
	return nil
}

// FloatingIPDiscard subcommand function
func FloatingIPDiscard(c *cli.Context) error {
	debugCmdFuncInfo(c)
	floatingIPSvc, formatter := WireUpFloatingIP(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := floatingIPSvc.DiscardFloatingIP(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discard floating IP", err)
	}
	return nil
}
