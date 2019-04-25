package cmd

import (
	"fmt"
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

	labelables := make([]types.Labelable, len(scripts))
	for i := 0; i < len(scripts); i++ {
		labelables[i] = types.Labelable(scripts[i])
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	scripts = make([]*types.Script, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Script)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Script, got a %T", labelable))
		}
		scripts[i] = s
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

	_, labelNamesByID := LabelLoadsMapping(c)
	script.FillInLabelNames(labelNamesByID)
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

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
		scriptIn["label_ids"] = labelsIdsArr
	}

	script, err := scriptSvc.CreateScript(&scriptIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create script", err)
	}

	script.FillInLabelNames(labelNamesByID)
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
	scriptIn := map[string]interface{}{}
	if c.String("name") != "" {
		scriptIn["name"] = strings.Split(c.String("name"), ",")
	}
	if c.String("description") != "" {
		scriptIn["description"] = strings.Split(c.String("description"), ",")
	}
	if c.String("code") != "" {
		scriptIn["code"] = strings.Split(c.String("code"), ",")
	}
	if c.String("parameters") != "" {
		scriptIn["parameters"] = strings.Split(c.String("parameters"), ",")
	}

	script, err := scriptSvc.UpdateScript(&scriptIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update script", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	script.FillInLabelNames(labelNamesByID)
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
