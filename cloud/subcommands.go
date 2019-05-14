package cloud

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cloud/cloud_providers"
	"github.com/ingrammicro/concerto/cloud/generic_images"
	"github.com/ingrammicro/concerto/cloud/server_plan"
	"github.com/ingrammicro/concerto/cloud/servers"
	"github.com/ingrammicro/concerto/cloud/ssh_profiles"
)

// SubCommands returns cloud commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "servers",
			Usage:       "Provides information on servers",
			Subcommands: append(servers.SubCommands()),
		},
		{
			Name:        "generic-images",
			Usage:       "Provides information on generic images",
			Subcommands: append(generic_images.SubCommands()),
		},
		{
			Name:        "ssh-profiles",
			Usage:       "Provides information on SSH profiles",
			Subcommands: append(ssh_profiles.SubCommands()),
		},
		{
			Name:        "cloud-providers",
			Usage:       "Provides information on cloud providers",
			Subcommands: append(cloud_providers.SubCommands()),
		},
		{
			Name:        "server-plans",
			Usage:       "Provides information on server plans",
			Subcommands: append(server_plan.SubCommands()),
		},
	}
}
