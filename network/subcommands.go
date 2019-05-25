package network

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/network/firewall_profiles"
	"github.com/ingrammicro/concerto/network/floating_ips"
)

// SubCommands returns network commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "firewall-profiles",
			Usage:       "Provides information about firewall profiles",
			Subcommands: append(firewall_profiles.SubCommands()),
		},
		{
			Name:        "floating-ips",
			Usage:       "Provides information about floating IPs",
			Subcommands: append(floating_ips.SubCommands()),
		},
	}
}
