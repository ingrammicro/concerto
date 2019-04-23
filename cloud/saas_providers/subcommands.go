package saas_providers

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns saas providers commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the SaaS providers supported by the platform.",
			Action: cmd.SaasProviderList,
		},
	}
}
