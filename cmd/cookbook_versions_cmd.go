package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
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

	// creates new cookbook_version
	cookbookVersion, err := svc.CreateCookbookVersion(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create cookbook version data", err)
	}

	// uploads new cookbook_version file
	err = svc.UploadCookbookVersion(sourceFilePath, cookbookVersion.UploadURL)
	if err != nil {
		formatter.PrintFatal("Couldn't upload cookbook version data", err)
	}

	// processes the new cookbook_version
	cookbookVersion, err = svc.ProcessCookbookVersion(utils.FlagConvertParams(c), cookbookVersion.ID)
	if err != nil {
		formatter.PrintFatal("Couldn't process cookbook version", err)
	}

	if err = formatter.PrintItem(*cookbookVersion); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}

	return nil
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
