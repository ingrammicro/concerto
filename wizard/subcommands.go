package wizard

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/wizard/apps"
	"github.com/ingrammicro/concerto/wizard/cloud_providers"
	"github.com/ingrammicro/concerto/wizard/locations"
	"github.com/ingrammicro/concerto/wizard/server_plans"
)

// SubCommands returns wizard commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "apps",
			Usage:       "Provides information about apps",
			Subcommands: append(apps.SubCommands()),
		},
		{
			Name:        "cloud-providers",
			Usage:       "Provides information about cloud providers",
			Subcommands: append(cloud_providers.SubCommands()),
		},
		{
			Name:        "locations",
			Usage:       "Provides information about locations",
			Subcommands: append(locations.SubCommands()),
		},
		{
			Name:        "server-plans",
			Usage:       "Provides information about server plans",
			Subcommands: append(server_plans.SubCommands()),
		},
	}
}
