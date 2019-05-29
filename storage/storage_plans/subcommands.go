package storage_plans

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns storage plans commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "show",
			Usage:  "Shows information about the storage plan identified by the given id",
			Action: cmd.StoragePlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Storage Plan Id",
				},
			},
		},
	}
}
