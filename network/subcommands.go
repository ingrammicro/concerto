package network

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/network/firewall_profiles"
)

// SubCommands returns network commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "firewall-profiles",
			Usage:       "Provides information about firewall profiles",
			Subcommands: append(firewall_profiles.SubCommands()),
		},
	}
}
