package audit

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands return Event subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list_events",
			Usage:  "Returns information about the events related to the account group.",
			Action: cmd.EventList,
		},
		{
			Name:   "list_system_events",
			Usage:  "Returns information about system-wide events.",
			Action: cmd.SysEventList,
		},
	}
}
