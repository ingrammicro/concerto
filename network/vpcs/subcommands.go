package vpcs

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns VPCs commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing VPCs",
			Action: cmd.VPCList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the VPC identified by the given id",
			Action: cmd.VPCShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "VPC Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new VPC",
			Action: cmd.VPCCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the VPC",
				},
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR of the VPC",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which the VPC is",
				},
				cli.StringFlag{
					Name:  "realm-provider-name",
					Usage: "Name of the provider realm in which the VPC is.",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with VPC",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing VPC identified by the given id",
			Action: cmd.VPCUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "VPC Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the VPC",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a VPC",
			Action: cmd.VPCDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "VPC Id",
				},
			},
		},
		{
			Name:   "add-label",
			Usage:  "This action assigns a single label from a single labelable resource",
			Action: cmd.LabelAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "VPC Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "vpc",
					Hidden: true,
				},
			},
		},
		{
			Name:   "remove-label",
			Usage:  "This action unassigns a single label from a single labelable resource",
			Action: cmd.LabelRemove,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "VPC Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "vpc",
					Hidden: true,
				},
			},
		},
	}
}
