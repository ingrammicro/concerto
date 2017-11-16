package cmdpolling

import (
	"github.com/codegangsta/cli"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "register",
			Usage:  "Registers concerto agent within an imported Host",
			Action: cmdRegister,
		},
		{
			Name:   "start",
			Usage:  "Starts a polling routine to check and execute pending scripts",
			Action: cmdStart,
		},
		{
			Name:   "stop",
			Usage:  "Stops the running polling process",
			Action: cmdStop,
		},
	}
}
