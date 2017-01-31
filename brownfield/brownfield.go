package brownfield

import "github.com/codegangsta/cli"

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "configure",
			Usage:  "Configures concerto agent within an imported brownfield Host",
			Action: cmdConfigure,
		},
	}
}
