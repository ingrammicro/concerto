package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
	"strings"
)

// WireUpTemplate prepares common resources to send request to Concerto API
func WireUpTemplate(c *cli.Context) (ts *blueprint.TemplateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ts, err = blueprint.NewTemplateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up template service", err)
	}

	return ts, f
}

// TemplateList subcommand function
func TemplateList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	templates, err := templateSvc.GetTemplateList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}

	labelables := make([]types.Labelable, len(templates))
	for i := 0; i < len(templates); i++ {
		labelables[i] = types.Labelable(&templates[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	templates = make([]types.Template, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		tpl, ok := labelable.(*types.Template)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Template, got a %T", labelable))
		}
		templates[i] = *tpl
	}

	if err = formatter.PrintList(templates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateShow subcommand function
func TemplateShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	template, err := templateSvc.GetTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateCreate subcommand function
func TemplateCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"name", "generic_image_id"}, formatter)
	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"cookbook_versions", "configuration_attributes"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateIn := map[string]interface{}{
		"name":                     c.String("name"),
		"generic_image_id":         c.String("generic_image_id"),
	}

	if c.IsSet("run_list") {
		templateIn["run_list"] = utils.RemoveDuplicates(strings.Split(c.String("run_list"), ","))
	}
	if c.IsSet("cookbook_versions") {
		templateIn["cookbook_versions"] = (*params)["cookbook_versions"]
	}
	if c.IsSet("configuration_attributes") {
		templateIn["configuration_attributes"] = (*params)["configuration_attributes"]
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
		templateIn["label_ids"] = labelsIdsArr
	}

	template, err := templateSvc.CreateTemplate(&templateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create template", err)
	}

	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateUpdate subcommand function
func TemplateUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"cookbook_versions", "configuration_attributes"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateIn := map[string]interface{}{}
	if c.IsSet("name") {
		templateIn["name"] = c.String("name")
	}
	if c.IsSet("run_list") {
		templateIn["run_list"] = utils.RemoveDuplicates(strings.Split(c.String("run_list"), ","))
	}
	if c.IsSet("cookbook_versions") {
		templateIn["cookbook_versions"] = (*params)["cookbook_versions"]
	}
	if c.IsSet("configuration_attributes") {
		templateIn["configuration_attributes"] = (*params)["configuration_attributes"]
	}

	template, err := templateSvc.UpdateTemplate(&templateIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update template", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateCompile subcommand function
func TemplateCompile(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	template, err := templateSvc.CompileTemplate(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't compile template", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateDelete subcommand function
func TemplateDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := templateSvc.DeleteTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete template", err)
	}
	return nil
}

// =========== Template Scripts =============

// TemplateScriptList subcommand function
func TemplateScriptList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type"}, formatter)
	templateScripts, err := templateScriptSvc.GetTemplateScriptList(c.String("template_id"), c.String("type"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintList(*templateScripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptShow subcommand function
func TemplateScriptShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)
	templateScript, err := templateScriptSvc.GetTemplateScript(c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptCreate subcommand function
func TemplateScriptCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type", "script_id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"parameter_values"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateScript, err := templateScriptSvc.CreateTemplateScript(params, c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't create templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptUpdate subcommand function
func TemplateScriptUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	// TODO si necessary: type script_id parameter_values ?
	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"parameter_values"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateScript, err := templateScriptSvc.UpdateTemplateScript(params, c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptDelete subcommand function
func TemplateScriptDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)
	err := templateScriptSvc.DeleteTemplateScript(c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete templateScript", err)
	}
	return nil
}

// TemplateScriptReorder subcommand function
func TemplateScriptReorder(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type", "script_ids"}, formatter)
	params, err := utils.FlagConvertParamsJSON(c, []string{"script_ids"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateScript, err := templateScriptSvc.ReorderTemplateScript(params, c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't reorder templateScript", err)
	}
	if err = formatter.PrintList(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// =========== Template Servers =============

// TemplateServersList subcommand function
func TemplateServersList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id"}, formatter)
	templateServers, err := templateSvc.GetTemplateServerList(c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template servers data", err)
	}
	if err = formatter.PrintList(*templateServers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
