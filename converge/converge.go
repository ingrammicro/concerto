package converge

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"path"
	"regexp"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
)

//CmbConverge Converges Host to original Blueprint
func CmbConverge(c *cli.Context) error {

	var firstBootJSONChef string

	if runtime.GOOS == "windows" {
		firstBootJSONChef = path.Join("c:\\chef", "first-boot.json")
	} else {
		firstBootJSONChef = path.Join("/etc/chef", "first-boot.json")
	}

	if utils.FileExists(firstBootJSONChef) {
		garbageOutput, _ := regexp.Compile("[\\[][^\\[|^\\]]*[\\]]\\s[A-Z]*:\\s")
		output, _ := regexp.Compile("Chef Run")
		cmd := exec.Command("chef-client", "-j", firstBootJSONChef)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Errorf("%s", err.Error())
		}
		ls := bufio.NewReader(stdout)
		err = cmd.Start()
		if err != nil {
			log.Errorf("%s", err.Error())
		}

		x := 0

		for {
			line, isPrefix, err := ls.ReadLine()
			if isPrefix {
				log.Errorf("%s", errors.New("isPrefix: true"))
			}
			if err != nil {
				if err != io.EOF {
					log.Errorf("%s", err.Error())
				}
				break
			}
			x = x + 1
			outputLine := garbageOutput.ReplaceAllString(string(line), "")
			if output.MatchString(outputLine) {
				log.Infof("%s", outputLine)
			} else {
				log.Debugf("%s", outputLine)
			}

		}
		err = cmd.Wait()
		if err != nil {
			log.Errorf("%s", err.Error())
		}
	} else {
		log.Fatalf("Make sure %s chef client configuration exists.", firstBootJSONChef)
	}
	return nil
}
