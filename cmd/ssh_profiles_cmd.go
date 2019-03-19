package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpSSHProfile prepares common resources to send request to Concerto API
func WireUpSSHProfile(c *cli.Context) (ds *cloud.SSHProfileService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewSSHProfileService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up sshProfile service", err)
	}

	return ds, f
}

// SSHProfileList subcommand function
func SSHProfileList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	sshProfiles, err := sshProfileSvc.GetSSHProfileList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive sshProfile data", err)
	}

	labelables := make([]types.Labelable, len(sshProfiles))
	for i:=0; i< len(sshProfiles); i++ {
		labelables[i] = types.Labelable(&sshProfiles[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)
	sshProfiles = make([]types.SSHProfile, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		sshP, ok := labelable.(*types.SSHProfile)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.SSHProfile, got a %T", labelable))
		}
		sshProfiles[i] = *sshP
	}

	if err = formatter.PrintList(sshProfiles); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileShow subcommand function
func SSHProfileShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	sshProfile, err := sshProfileSvc.GetSSHProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive sshProfile data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	sshProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileCreate subcommand function
func SSHProfileCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"name", "public_key"}, formatter)
	sshProfileIn := map[string]interface{}{
		"name":       c.String("name"),
		"public_key": c.String("public_key"),
	}
	if c.String("private_key") != "" {
		sshProfileIn["private_key"] = c.String("private_key")
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		labelsIdsArr := LabelResolution(c, c.String("labels"), labelIDsByName)
		sshProfileIn["label_ids"] = labelsIdsArr
	}

	sshProfile, err := sshProfileSvc.CreateSSHProfile(&sshProfileIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create sshProfile", err)
	}

	sshProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileUpdate subcommand function
func SSHProfileUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	sshProfile, err := sshProfileSvc.UpdateSSHProfile(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update sshProfile", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	sshProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileDelete subcommand function
func SSHProfileDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := sshProfileSvc.DeleteSSHProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete sshProfile", err)
	}
	return nil
}
