package dispatcher

import (
	"github.com/codegangsta/cli"
)

// SubCommands returns dispatcher commands
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
