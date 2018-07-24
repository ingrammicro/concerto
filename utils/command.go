package utils

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// TODO remove after migration

//FlagsRequired checks for required flags, and show usage if requirements not met
func FlagsRequired(c *cli.Context, flags []string) {
	parameters := false
	for _, flag := range flags {
		if !c.IsSet(flag) {
			log.Warn(fmt.Sprintf("Please use parameter --%s", flag))
			parameters = true
		}
	}

	if parameters {
		fmt.Printf("\n")
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(2)
	}
}
