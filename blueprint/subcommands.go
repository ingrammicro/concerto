package blueprint

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/blueprint/cookbook_versions"
	"github.com/ingrammicro/concerto/blueprint/scripts"
	"github.com/ingrammicro/concerto/blueprint/templates"
)

// SubCommands returns blueprint commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "cookbook-versions",
			Usage:       "Provides information on chef cookbook versions",
			Subcommands: append(cookbook_versions.SubCommands()),
		},
		{
			Name:        "scripts",
			Usage:       "Allow the user to manage the scripts they want to run on the servers",
			Subcommands: append(scripts.SubCommands()),
		},
		{
			Name:        "templates",
			Usage:       "Provides information on templates",
			Subcommands: append(templates.SubCommands()),
		},
	}
}
