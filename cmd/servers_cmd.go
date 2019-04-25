package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpServer prepares common resources to send request to Concerto API
func WireUpServer(c *cli.Context) (ds *cloud.ServerService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewServerService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up server service", err)
	}

	return ds, f
}

// ServerList subcommand function
func ServerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	servers, err := serverSvc.GetServerList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}

	labelables := make([]types.Labelable, len(servers))
	for i := 0; i < len(servers); i++ {
		labelables[i] = types.Labelable(servers[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	servers = make([]*types.Server, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Server)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Server, got a %T", labelable))
		}
		servers[i] = s
	}
	if err = formatter.PrintList(servers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerShow subcommand function
func ServerShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.GetServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerCreate subcommand function
func ServerCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"name", "ssh-profile-id", "firewall-profile-id", "template-id", "server-plan-id", "cloud-account-id"}, formatter)
	serverIn := map[string]interface{}{
		"name":                c.String("name"),
		"ssh_profile_id":      c.String("ssh-profile-id"),
		"firewall_profile_id": c.String("firewall-profile-id"),
		"template_id":         c.String("template-id"),
		"server_plan_id":      c.String("server-plan-id"),
		"cloud_account_id":    c.String("cloud-account-id"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
		serverIn["label_ids"] = labelsIdsArr
	}

	server, err := serverSvc.CreateServer(&serverIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create server", err)
	}

	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerUpdate subcommand function
func ServerUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.UpdateServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerBoot subcommand function
func ServerBoot(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.BootServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't boot server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerReboot subcommand function
func ServerReboot(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.RebootServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't reboot server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerShutdown subcommand function
func ServerShutdown(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.ShutdownServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't shutdown server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerOverride subcommand function
func ServerOverride(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.OverrideServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't override server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerDelete subcommand function
func ServerDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := serverSvc.DeleteServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete server", err)
	}
	return nil
}

// ========= Events ========

// EventsList subcommand function
func EventsList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	events, err := svc.GetEventsList(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

//======= Operational Scripts ==========

// OperationalScriptsList subcommand function
func OperationalScriptsList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	scripts, err := svc.GetOperationalScriptsList(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}
	if err = formatter.PrintList(scripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// OperationalScriptExecute subcommand function
func OperationalScriptExecute(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"server-id", "script-id"}, formatter)
	in := &map[string]interface{}{}
	scriptOut, err := serverSvc.ExecuteOperationalScript(in, c.String("server-id"), c.String("script-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't execute operational script", err)
	}
	if err = formatter.PrintItem(*scriptOut); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
