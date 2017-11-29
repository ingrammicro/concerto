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
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "time, t",
					Usage: "Polling ping time interval (seconds)",
					Value: 30,
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops the running polling process",
			Action: cmdStop,
		},
	}
}
