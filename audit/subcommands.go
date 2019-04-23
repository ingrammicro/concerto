package audit

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns event commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list-events",
			Usage:  "Returns information about the events related to the account group.",
			Action: cmd.EventList,
		},
		{
			Name:   "list-system-events",
			Usage:  "Returns information about system-wide events.",
			Action: cmd.SysEventList,
		},
	}
}
