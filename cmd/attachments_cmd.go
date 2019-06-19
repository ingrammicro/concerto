package cmd

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpAttachment prepares common resources to send request to Concerto API
func WireUpAttachment(c *cli.Context) (scs *blueprint.AttachmentService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	scs, err = blueprint.NewAttachmentService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up attachment service", err)
	}

	return scs, f
}

// AttachmentShow subcommand function
func AttachmentShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	attachmentSvc, formatter := WireUpAttachment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	attachment, err := attachmentSvc.GetAttachment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive attachment data", err)
	}

	if err = formatter.PrintItem(*attachment); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// AttachmentDownload subcommand function
func AttachmentDownload(c *cli.Context) error {
	debugCmdFuncInfo(c)
	attachmentSvc, formatter := WireUpAttachment(c)

	checkRequiredFlags(c, []string{"id", "filepath"}, formatter)
	attachment, err := attachmentSvc.GetAttachment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive attachment data", err)
	}

	realFileName, status, err := attachmentSvc.DownloadAttachment(attachment.DownloadURL, c.String("filepath"))
	if err == nil && status != 200 {
		err = fmt.Errorf("obtained non-ok response when downloading attachment %s", attachment.DownloadURL)
	}
	if err != nil {
		return err
	}
	logrus.Info("Available at:", realFileName)
	return nil
}

// AttachmentDelete subcommand function
func AttachmentDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	attachmentSvc, formatter := WireUpAttachment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := attachmentSvc.DeleteAttachment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete attachment", err)
	}
	return nil
}
