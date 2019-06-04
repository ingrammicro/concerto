package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/storage"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpVolume prepares common resources to send request to Concerto API
func WireUpVolume(c *cli.Context) (ds *storage.VolumeService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = storage.NewVolumeService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up volume service", err)
	}

	return ds, f
}

// VolumeList subcommand function
func VolumeList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	volumes, err := volumeSvc.GetVolumeList(c.String("server-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive volume data", err)
	}

	labelables := make([]types.Labelable, len(volumes))
	for i := 0; i < len(volumes); i++ {
		labelables[i] = types.Labelable(volumes[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	volumes = make([]*types.Volume, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Volume)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.Volume, got a %T", labelable))
		}
		volumes[i] = v
	}
	if err = formatter.PrintList(volumes); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VolumeShow subcommand function
func VolumeShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	volume, err := volumeSvc.GetVolume(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive volume data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VolumeCreate subcommand function
func VolumeCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"name", "size", "cloud-account-id", "storage-plan-id"}, formatter)

	volumeIn := map[string]interface{}{
		"name":             c.String("name"),
		"size":             c.Int("size"),
		"cloud_account_id": c.String("cloud-account-id"),
		"storage_plan_id":  c.String("storage-plan-id"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		volumeIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	volume, err := volumeSvc.CreateVolume(&volumeIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create volume", err)
	}

	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VolumeUpdate subcommand function
func VolumeUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)

	volumeIn := map[string]interface{}{
		"name": c.String("name"),
	}

	volume, err := volumeSvc.UpdateVolume(&volumeIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update volume", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	volume.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*volume); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VolumeAttach subcommand function
func VolumeAttach(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"id", "server-id"}, formatter)

	volumeIn := map[string]interface{}{
		"attached_server_id": c.String("server-id"),
	}

	server, err := volumeSvc.AttachVolume(&volumeIn, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't attach volume", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// VolumeDetach subcommand function
func VolumeDetach(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := volumeSvc.DetachVolume(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't detach volume", err)
	}
	return nil
}

// VolumeDelete subcommand function
func VolumeDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := volumeSvc.DeleteVolume(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete volume", err)
	}
	return nil
}

// VolumeDiscard subcommand function
func VolumeDiscard(c *cli.Context) error {
	debugCmdFuncInfo(c)
	volumeSvc, formatter := WireUpVolume(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := volumeSvc.DiscardVolume(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discard volume", err)
	}
	return nil
}
