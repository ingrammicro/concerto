package cmd

import (
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

	labelables := make([]*types.LabelableFields, 0, len(servers))
	for i := range servers {
		labelables = append(labelables, &servers[i].LabelableFields)
	}

	filteredLabelables := LabelFiltering(c, labelables)

	tmp := servers
	servers = nil
	if len(filteredLabelables) > 0 {
		for _, labelable := range filteredLabelables {
			for i := range tmp {
				if &tmp[i].LabelableFields == labelable {
					servers = append(servers, tmp[i])
				}
			}
		}
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

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerCreate subcommand function
func ServerCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"name", "ssh_profile_id", "firewall_profile_id", "template_id", "server_plan_id", "cloud_account_id"}, formatter)
	serverIn := map[string]interface{}{
		"name":                c.String("name"),
		"ssh_profile_id":      c.String("ssh_profile_id"),
		"firewall_profile_id": c.String("firewall_profile_id"),
		"template_id":         c.String("template_id"),
		"server_plan_id":      c.String("server_plan_id"),
		"cloud_account_id":    c.String("cloud_account_id"),
	}

	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"))
		serverIn["label_ids"] = labelsIdsArr
	}

	server, err := serverSvc.CreateServer(&serverIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create server", err)
	}

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
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

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
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

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
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

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
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

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
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

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&server.LabelableFields})
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

// ========= DNS ========
// DNSList subcommand function
func DNSList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	dnsSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	dnss, err := dnsSvc.GetDNSList(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive dns data", err)
	}
	if err = formatter.PrintList(dnss); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ========= Events ========

// EventsList subcommand function
func EventsList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	dnsSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	events, err := dnsSvc.GetEventsList(c.String("id"))
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
	dnsSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	scripts, err := dnsSvc.GetOperationalScriptsList(c.String("id"))
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

	checkRequiredFlags(c, []string{"server_id", "script_id"}, formatter)
	server, err := serverSvc.ExecuteOperationalScript(utils.FlagConvertParams(c), c.String("server_id"), c.String("script_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't execute operational script", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
