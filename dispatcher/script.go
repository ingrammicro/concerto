package dispatcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/webservice"
)

const (
	characterizationsEndpoint = "blueprint/script_characterizations?type=%s"
	characterizationEndpoint  = "blueprint/script_characterizations/%s"
	conclusionsEndpoint       = "blueprint/script_conclusions"
)

// ScriptCharacterization stores Script Characterization data
type ScriptCharacterization struct {
	Order      int               `json:"execution_order"`
	UUID       string            `json:"uuid"`
	Script     Script            `json:"script"`
	Parameters map[string]string `json:"parameter_values"`
}

// Script stores Script data
type Script struct {
	Code            string   `json:"code"`
	UUID            string   `json:"uuid"`
	AttachmentPaths []string `json:"attachment_paths"`
}

// ScriptConclusion stores Script Conclusion data
type ScriptConclusion struct {
	UUID       string `json:"script_characterization_id"`
	Output     string `json:"output"`
	ExitCode   int    `json:"exit_code"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}

// ScriptConclusionRoot stores Script Conclusion Root data
type ScriptConclusionRoot struct {
	Root ScriptConclusion `json:"script_conclusion"`
}

// SubCommands return Script subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "boot",
			Usage:  "Executes script characterizations associated to booting state of host",
			Action: cmdBoot,
		},
		{
			Name:   "operational",
			Usage:  "Executes all script characterizations associated to operational state of host or the one with the given id",
			Action: cmdOperational,
		},
		{
			Name:   "shutdown",
			Usage:  "Executes script characterizations associated to shutdown state of host",
			Action: cmdShutdown,
		},
	}
}

func executeScriptCharacterization(script ScriptCharacterization, directoryPath string) (conclusion ScriptConclusionRoot) {
	output, exitCode, startedAt, finishedAt := utils.ExecCode(script.Script.Code, directoryPath, script.Script.UUID)

	conclusion.Root.UUID = script.UUID
	conclusion.Root.Output = output
	conclusion.Root.ExitCode = exitCode
	conclusion.Root.StartedAt = startedAt.Format(utils.TimeStampLayout)
	conclusion.Root.FinishedAt = finishedAt.Format(utils.TimeStampLayout)

	return conclusion
}

func execute(phase string, scriptScharacterizationUUID string) {
	var scriptChars []ScriptCharacterization
	webservice, err := webservice.NewWebService()
	if err != nil {
		log.Fatal(err)
	}
	if scriptScharacterizationUUID == "" {
		log.Debugf("Current Script Characterization %s", phase)
		data, _, err := webservice.Get(fmt.Sprintf(characterizationsEndpoint, phase))
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf(string(data))

		err = json.Unmarshal(data, &scriptChars)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Debugf("%s Script Characterization %s", phase, scriptScharacterizationUUID)
		data, _, err := webservice.Get(fmt.Sprintf(characterizationEndpoint, scriptScharacterizationUUID))
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf(string(data))

		scriptChars = make([]ScriptCharacterization, 1)
		err = json.Unmarshal(data, &scriptChars[0])
		if err != nil {
			log.Fatal(err)
		}
	}
	scripts := ByOrder(scriptChars)

	for _, ex := range scripts {
		log.Infof("------------------------------------------------------------------------------------------------")
		path, err := ioutil.TempDir("", "concerto")
		if err != nil {
			log.Fatal(err)
		}

		os.Setenv("ATTACHMENT_DIR", fmt.Sprintf("%s/%s", path, "attachments"))

		log.Infof("UUID: %s", ex.UUID)
		log.Infof("Home Folder: %s", path)
		err = os.Mkdir(os.Getenv("ATTACHMENT_DIR"), 0777)
		if err != nil {
			log.Fatal(err)
		}

		// Seting up Environment Variables
		log.Infof("Environment Variables")
		for index, value := range ex.Parameters {
			os.Setenv(index, value)
			log.Infof("\t - %s=%s", index, value)
		}

		if len(ex.Script.AttachmentPaths) > 0 {
			log.Infof("Attachment Folder: %s", os.Getenv("ATTACHMENT_DIR"))
			// Downloading Attachements
			log.Infof("Attachments")
			if err != nil {
				log.Fatal(err)
			}
			for _, endpoint := range ex.Script.AttachmentPaths {
				filename, err := webservice.GetFile(endpoint, os.Getenv("ATTACHMENT_DIR"))
				if err != nil {
					log.Fatal(err)
				}
				log.Infof("\t - %s --> %s", endpoint, filename)
			}
		}

		json, err := json.Marshal(executeScriptCharacterization(ex, path))
		if err != nil {
			log.Fatal(err)
		}

		_, _, err = webservice.Post(conclusionsEndpoint, json)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("------------------------------------------------------------------------------------------------")
	}
}

func cmdBoot(c *cli.Context) error {
	execute("boot", "")
	return nil
}

func cmdOperational(c *cli.Context) error {
	execute("operational", c.Args().Get(0))
	return nil
}

func cmdShutdown(c *cli.Context) error {
	execute("shutdown", "")
	return nil
}
