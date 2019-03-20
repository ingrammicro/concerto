package labels

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the current labels existing in the platform for the user",
			Action: cmd.LabelList,
		},
	}
}
