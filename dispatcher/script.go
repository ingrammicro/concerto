package dispatcher

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"io/ioutil"
	"os"
)

func cmdBoot(c *cli.Context) error {
	execute(c, "boot", "")
	return nil
}

func cmdOperational(c *cli.Context) error {
	execute(c, "operational", c.Args().Get(0))
	return nil
}

func cmdShutdown(c *cli.Context) error {
	execute(c, "shutdown", "")
	return nil
}

func execute(c *cli.Context, phase string, scriptCharacterizationUUID string) {
	var scriptChars []*types.ScriptCharacterization
	dispatcherSvc, formatter := cmd.WireUpDispatcher(c)

	var err error
	log.Debugf("Current Script Characterization %s (UUID=%s)", phase, scriptCharacterizationUUID)
	if scriptCharacterizationUUID == "" {
		scriptChars, err = dispatcherSvc.GetDispatcherScriptCharacterizationsByType(phase)
	} else {
		scriptChars, err = dispatcherSvc.GetDispatcherScriptCharacterizationsByUUID(scriptCharacterizationUUID)
	}
	if err != nil {
		formatter.PrintFatal("Couldn't receive Script Characterization data", err)
	}

	for _, sc := range scriptChars {
		log.Infof("------------------------------------------------------------------------------------------------")
		path, err := ioutil.TempDir("", "cio")
		if err != nil {
			formatter.PrintFatal("Couldn't create temporary directory", err)
		}
		defer os.RemoveAll(path)

		if err = os.Setenv("ATTACHMENT_DIR", fmt.Sprintf("%s/%s", path, "attachments")); err != nil {
			formatter.PrintFatal("Couldn't set attachments directory as environment variable", err)
		}

		log.Infof("UUID: %s", sc.UUID)
		log.Infof("Home Folder: %s", path)

		attachmentDir := os.Getenv("ATTACHMENT_DIR")
		if err = os.Mkdir(attachmentDir, 0777); err != nil {
			formatter.PrintFatal("Couldn't create attachments directory", err)
		}

		// Setting up environment Variables
		log.Infof("Environment Variables")
		for index, value := range sc.Parameters {
			if err = os.Setenv(index, value); err != nil {
				formatter.PrintFatal(fmt.Sprintf("Couldn't set environment variable %s:%s", index, value), err)
			}
			log.Infof("\t - %s=%s", index, value)
		}

		if len(sc.Script.AttachmentPaths) > 0 {
			log.Infof("Attachment Folder: %s", attachmentDir)
			log.Infof("Attachments")
			for _, endpoint := range sc.Script.AttachmentPaths {
				realFileName, _, err := dispatcherSvc.DownloadAttachment(endpoint, attachmentDir)
				if err != nil {
					formatter.PrintFatal("Couldn't download attachment", err)
				}
				log.Infof("\t - %s --> %s", endpoint, realFileName)
			}
		}

		output, exitCode, startedAt, finishedAt := utils.ExecCode(sc.Script.Code, path, sc.Script.UUID)
		scriptConclusionIn := map[string]interface{}{
			"script_characterization_id": sc.UUID,
			"output":                     output,
			"exit_code":                  exitCode,
			"started_at":                 startedAt.Format(utils.TimeStampLayout),
			"finished_at":                finishedAt.Format(utils.TimeStampLayout),
		}
		scriptConclusionRootIn := map[string]interface{}{
			"script_conclusion": scriptConclusionIn,
		}
		_, _, err = dispatcherSvc.ReportScriptConclusions(&scriptConclusionRootIn)
		if err != nil {
			formatter.PrintFatal("Couldn't send script_conclusions report data", err)
		}
		log.Infof("------------------------------------------------------------------------------------------------")
	}
}
