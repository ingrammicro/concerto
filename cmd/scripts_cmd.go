package cmd

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpScript prepares common resources to send request to Concerto API
func WireUpScript(c *cli.Context) (scs *blueprint.ScriptService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	scs, err = blueprint.NewScriptService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up script service", err)
	}

	return scs, f
}

// ScriptsList subcommand function
func ScriptsList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	scripts, err := scriptSvc.GetScriptList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}

	labelables := make([]*types.LabelableFields, 0, len(scripts))
	for i := range scripts {
		labelables = append(labelables, &scripts[i].LabelableFields)
	}

	filteredLabelables := LabelFiltering(c, labelables)

	tmp := scripts
	scripts = nil
	if len(filteredLabelables) > 0 {
		for _, labelable := range filteredLabelables {
			for i := range tmp {
				if &tmp[i].LabelableFields == labelable {
					scripts = append(scripts, tmp[i])
				}
			}
		}
	}

	if err = formatter.PrintList(scripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ScriptShow subcommand function
func ScriptShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	script, err := scriptSvc.GetScript(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&script.LabelableFields})
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ScriptCreate subcommand function
func ScriptCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"name", "description", "code"}, formatter)
	scriptIn := map[string]interface{}{
		"name":        c.String("name"),
		"description": c.String("description"),
		"code":        c.String("code"),
	}
	if c.String("parameters") != "" {
		scriptIn["parameters"] = strings.Split(c.String("parameters"), ",")
	}
	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"))
		scriptIn["label_ids"] = labelsIdsArr
	}

	script, err := scriptSvc.CreateScript(&scriptIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create script", err)
	}

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&script.LabelableFields})
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ScriptUpdate subcommand function
func ScriptUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	script, err := scriptSvc.UpdateScript(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update script", err)
	}

	LabelAssignNamesForIDs(c, []*types.LabelableFields{&script.LabelableFields})
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ScriptDelete subcommand function
func ScriptDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := scriptSvc.DeleteScript(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete script", err)
	}
	return nil
}
