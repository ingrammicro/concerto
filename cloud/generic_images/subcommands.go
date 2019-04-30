package generic_images

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns generic images commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "This action lists the available generic images.",
			Action: cmd.GenericImageList,
		},
	}
}
