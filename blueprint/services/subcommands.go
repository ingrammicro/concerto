package services

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands return Services subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available services",
			Action: cmd.ServiceList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific service",
			Action: cmd.ServiceShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Service Id",
				},
			},
		},
	}
}
