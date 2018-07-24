package cloud_providers

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands return Cloud Provider subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available cloud providers.",
			Action: cmd.CloudProviderList,
		},
	}
}
