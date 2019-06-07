package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpCookbookVersion prepares common resources to send request to Concerto API
func WireUpCookbookVersion(c *cli.Context) (sv *blueprint.CookbookVersionService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	sv, err = blueprint.NewCookbookVersionService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up CookbookVersion service", err)
	}

	return sv, f
}

// CookbookVersionList subcommand function
func CookbookVersionList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCookbookVersion(c)

	cookbookVersions, err := svc.GetCookbookVersionList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cookbook versions data", err)
	}

	labelables := make([]types.Labelable, len(cookbookVersions))
	for i := 0; i < len(cookbookVersions); i++ {
		labelables[i] = types.Labelable(cookbookVersions[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)
	cookbookVersions = make([]*types.CookbookVersion, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		cb, ok := labelable.(*types.CookbookVersion)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.CookbookVersion, got a %T", labelable))
		}
		cookbookVersions[i] = cb
	}

	if err = formatter.PrintList(cookbookVersions); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CookbookVersionShow subcommand function
func CookbookVersionShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCookbookVersion(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cookbookVersion, err := svc.GetCookbookVersion(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cookbook version data", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cookbookVersion.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cookbookVersion); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CookbookVersionUpload subcommand function
// create/upload/process
func CookbookVersionUpload(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCookbookVersion(c)

	checkRequiredFlags(c, []string{"filepath"}, formatter)
	sourceFilePath := c.String("filepath")

	if !utils.FileExists(sourceFilePath) {
		formatter.PrintFatal("Invalid file path", fmt.Errorf("no such file or directory: %s", sourceFilePath))
	}

	cbIn := map[string]interface{}{}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		cbIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	// creates new cookbook_version
	cookbookVersion, err := svc.CreateCookbookVersion(&cbIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create cookbook version data", err)
	}

	// uploads new cookbook_version file
	err = svc.UploadCookbookVersion(sourceFilePath, cookbookVersion.UploadURL)
	if err != nil {
		cleanCookbookVersion(c, cookbookVersion.ID)
		formatter.PrintFatal("Couldn't upload cookbook version data", err)
	}

	// processes the new cookbook_version
	cookbookVersionID := cookbookVersion.ID
	cookbookVersion, err = svc.ProcessCookbookVersion(utils.FlagConvertParams(c), cookbookVersion.ID)
	if err != nil {
		cleanCookbookVersion(c, cookbookVersionID)
		formatter.PrintFatal("Couldn't process cookbook version", err)
	}

	cookbookVersion.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cookbookVersion); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}

	return nil
}

// cleanCookbookVersion deletes CookbookVersion. Ideally for cleaning at uploading error cases
func cleanCookbookVersion(c *cli.Context, cookbookVersionID string) {
	svc, formatter := WireUpCookbookVersion(c)
	if err := svc.DeleteCookbookVersion(cookbookVersionID); err != nil {
		formatter.PrintError("Couldn't clean failed cookbook version", err)
	}
}

// CookbookVersionDelete subcommand function
func CookbookVersionDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCookbookVersion(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := svc.DeleteCookbookVersion(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cookbook version", err)
	}
	return nil
}
