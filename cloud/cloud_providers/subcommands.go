package cloud_providers

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns cloud providers commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available cloud providers",
			Action: cmd.CloudProviderList,
		},
		{
			Name:   "list-storage-plans",
			Usage:  "This action lists the storage plans offered by the cloud provider identified by the given id",
			Action: cmd.CloudProviderStoragePlansList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud-provider-id",
					Usage: "Cloud provider id",
				},
			},
		},
	}
}
