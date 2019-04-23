package settings

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/settings/cloud_accounts"
)

// SubCommands returns settings commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "cloud-accounts",
			Usage:       "Provides information about cloud accounts",
			Subcommands: append(cloud_accounts.SubCommands()),
		},
	}
}
