package cmd

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
	"os"
	"regexp"
	"strings"
)

var templateCookbookVersionValueRegexp = regexp.MustCompile(`^(.*?)(~>|=|>=|<=|>|<|:)(\d+(?:\.\d+)+)$`)

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
		labelables[i] = types.Labelable(templates[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	templates = make([]*types.Template, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		tpl, ok := labelable.(*types.Template)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Template, got a %T", labelable))
		}
		templates[i] = tpl
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

	checkRequiredFlags(c, []string{"name", "generic-image-id"}, formatter)

	if c.IsSet("configuration-attributes") && c.IsSet("configuration-attributes-from-file") {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'configuration-attributes' or 'configuration-attributes-from-file'")
	}

	templateIn := map[string]interface{}{
		"name":             c.String("name"),
		"generic_image_id": c.String("generic-image-id"),
	}
	if c.IsSet("configuration-attributes-from-file") {
		caIn, err := convertFlagParamsToConfigurationAttributesFromFile(c, c.String("configuration-attributes-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input configuration attributes", err)
		}
		templateIn["configuration_attributes"] = caIn
	}
	if c.IsSet("configuration-attributes") {
		params, err := utils.FlagConvertParamsJSON(c, []string{"configuration-attributes"})
		if err != nil {
			formatter.PrintFatal("Cannot parse input configuration attributes", err)
		}
		templateIn["configuration_attributes"] = (*params)["configuration-attributes"]
	}
	if c.IsSet("run-list") {
		templateIn["run_list"] = utils.RemoveDuplicates(strings.Split(c.String("run-list"), ","))
	}
	if c.IsSet("cookbook-versions") {
		cbIn, err := convertFlagParamsToCookbookVersions(c, c.String("cookbook-versions"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input cookbook versions", err)
		}
		templateIn["cookbook_versions"] = cbIn
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

	if c.IsSet("configuration-attributes") && c.IsSet("configuration-attributes-from-file") {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'configuration-attributes' or 'configuration-attributes-from-file'")
	}

	templateIn := map[string]interface{}{}
	if c.IsSet("name") {
		templateIn["name"] = c.String("name")
	}
	if c.IsSet("configuration-attributes-from-file") {
		caIn, err := convertFlagParamsToConfigurationAttributesFromFile(c, c.String("configuration-attributes-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input configuration attributes", err)
		}
		templateIn["configuration_attributes"] = caIn
	}
	if c.IsSet("configuration-attributes") {
		params, err := utils.FlagConvertParamsJSON(c, []string{"configuration-attributes"})
		if err != nil {
			formatter.PrintFatal("Cannot parse input configuration attributes", err)
		}
		templateIn["configuration_attributes"] = (*params)["configuration-attributes"]
	}
	if c.IsSet("run-list") {
		templateIn["run_list"] = utils.RemoveDuplicates(strings.Split(c.String("run-list"), ","))
	}
	if c.IsSet("cookbook-versions") {
		cbIn, err := convertFlagParamsToCookbookVersions(c, c.String("cookbook-versions"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input cookbook versions", err)
		}
		templateIn["cookbook_versions"] = cbIn
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

	checkRequiredFlags(c, []string{"template-id", "type"}, formatter)
	templateScripts, err := templateScriptSvc.GetTemplateScriptList(c.String("template-id"), c.String("type"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintList(templateScripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptShow subcommand function
func TemplateScriptShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template-id"}, formatter)
	templateScript, err := templateScriptSvc.GetTemplateScript(c.String("template-id"), c.String("id"))
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

	checkRequiredFlags(c, []string{"template-id", "type", "script-id"}, formatter)

	templateScriptIn := map[string]interface{}{
		"type":      c.String("type"),
		"script_id": c.String("script-id"),
	}
	if c.IsSet("parameter-values") {
		tsIn, err := convertFlagParamsToParameterValues(c, c.String("parameter-values"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input parameter values", err)
		}
		templateScriptIn["parameter_values"] = tsIn
	}

	templateScript, err := templateScriptSvc.CreateTemplateScript(&templateScriptIn, c.String("template-id"))
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

	checkRequiredFlags(c, []string{"template-id", "id"}, formatter)

	templateScriptIn := map[string]interface{}{}
	if c.IsSet("parameter-values") {
		tsIn, err := convertFlagParamsToParameterValues(c, c.String("parameter-values"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input parameter values", err)
		}
		templateScriptIn["parameter_values"] = tsIn
	}

	templateScript, err := templateScriptSvc.UpdateTemplateScript(&templateScriptIn, c.String("template-id"), c.String("id"))
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

	checkRequiredFlags(c, []string{"id", "template-id"}, formatter)
	err := templateScriptSvc.DeleteTemplateScript(c.String("template-id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete templateScript", err)
	}
	return nil
}

// TemplateScriptReorder subcommand function
func TemplateScriptReorder(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template-id", "type", "script-ids"}, formatter)
	templateScriptIn := map[string]interface{}{
		"type":       c.String("type"),
		"script_ids": utils.RemoveDuplicates(strings.Split(c.String("script-ids"), ",")),
	}

	templateScript, err := templateScriptSvc.ReorderTemplateScript(&templateScriptIn, c.String("template-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't reorder templateScript", err)
	}
	if err = formatter.PrintList(templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// =========== Template Servers =============

// TemplateServersList subcommand function
func TemplateServersList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template-id"}, formatter)
	templateServers, err := templateSvc.GetTemplateServerList(c.String("template-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template servers data", err)
	}
	if err = formatter.PrintList(templateServers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// =========== Template helpers =============

// convertFlagParamsToCookbookVersions returns the json representation for the given friendly input format of cookbook versions assignation
// i.e: "wordpress:0.1.0,nano=2.0.1,1password~>1.3.0"
func convertFlagParamsToCookbookVersions(c *cli.Context, cbvsIn string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	cookbookVersions := make([]*types.CookbookVersion, 0)
	for _, cbvIn := range strings.Split(cbvsIn, ",") {
		values := templateCookbookVersionValueRegexp.FindStringSubmatch(cbvIn)
		if len(values) == 0 {
			return nil, fmt.Errorf("invalid input cookbook version format %s", cbvIn)
		}
		name, operator, version := values[1], values[2], values[3]
		if _, found := result[name]; found {
			return nil, fmt.Errorf("detected duplicated cookbook version name: %s", name)
		}

		// uploaded. It requires to map adequate version_id
		if operator == ":" {
			if len(cookbookVersions) == 0 {
				// data is loaded only once
				svc, formatter := WireUpCookbookVersion(c)
				cbvs, err := svc.GetCookbookVersionList()
				if err != nil {
					formatter.PrintFatal("cannot receive uploaded cookbook versions data", err)
				}
				cookbookVersions = cbvs
			}
			for _, cbv := range cookbookVersions {
				if name == cbv.Name && version == cbv.Version {
					result[name] = map[string]interface{}{"version_id": cbv.ID}
				}
			}
			// provided cookbook version does not match the available uploaded
			if _, found := result[name]; !found {
				return nil, fmt.Errorf("invalid cookbook version: %s does not match any of the cookbook versions uploaded to the platform", cbvIn)
			}
		} else {
			//supermarket
			result[name] = map[string]interface{}{"version": operator + " " + version} // at any case, it should leave a space between operator and version
		}
	}
	return result, nil
}

// convertFlagParamsToParameterValues returns the json representation for the given friendly input format of parameter-values assignation
func convertFlagParamsToParameterValues(c *cli.Context, cbIn string) (map[string]string, error) {
	result := map[string]string{}
	for _, cb := range strings.Split(cbIn, ",") {
		values := regexp.MustCompile(`(.*?)(:)(\w.*)`).FindStringSubmatch(cb)
		if len(values) == 0 {
			return nil, fmt.Errorf("invalid input parameter values format %s", cb)
		}
		name, _, value := values[1], values[2], values[3]
		if _, found := result[name]; found {
			return nil, fmt.Errorf("detected duplicated parameter name: %s", name)
		}
		result[name] = value
	}
	return result, nil
}

// convertFlagParamsToConfigurationAttributesFromFile returns the json representation of configuration attributes taken from the input file or STDIN
func convertFlagParamsToConfigurationAttributesFromFile(c *cli.Context, casIn string) (map[string]interface{}, error) {
	var content map[string]interface{}
	if casIn == "-" {
		// read from STDIN
		log.Info("Please, write configuration parameters json formatted:")
		if err := json.NewDecoder(os.Stdin).Decode(&content); err != nil {
			return nil, fmt.Errorf("invalid json formatted attributes")
		}
	} else {
		// read from file
		sourceFilePath := casIn
		if !utils.FileExists(sourceFilePath) {
			return nil, fmt.Errorf("invalid file path, no such file: %s", sourceFilePath)
		}

		attrsFile, err := os.Open(sourceFilePath)
		if err != nil {
			return nil, fmt.Errorf("cannot open file %s: %v", sourceFilePath, err)
		}
		defer attrsFile.Close()

		if err = json.NewDecoder(attrsFile).Decode(&content); err != nil {
			return nil, fmt.Errorf("invalid json formatted attributes in file %s", sourceFilePath)
		}
	}
	return content, nil
}
