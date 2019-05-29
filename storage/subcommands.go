package storage

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/storage/storage_plans"
	"github.com/ingrammicro/concerto/storage/volumes"
)

// SubCommands returns storage commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "plans",
			Usage:       "Provides information on storage plans",
			Subcommands: append(storage_plans.SubCommands()),
		},
		{
			Name:        "volumes",
			Usage:       "Provides information on storage volumes",
			Subcommands: append(volumes.SubCommands()),
		},
	}
}