package subnets

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns Subnets commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all Subnets of a VPC",
			Action: cmd.SubnetList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "VPC Id",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the Subnet identified by the given id",
			Action: cmd.SubnetShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Subnet Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new Subnet inside the specified VPC",
			Action: cmd.SubnetCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "VPC Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the Subnet",
				},
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR of the Subnet",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Type of the Subnet (among 'only_public', 'only_private' and 'public_and_private')",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing Subnet identified by the given id",
			Action: cmd.SubnetUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Subnet Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the Subnet",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a Subnet",
			Action: cmd.SubnetDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Subnet Id",
				},
			},
		},
		{
			Name:   "list-servers",
			Usage:  "Lists servers belonging to the subnet identified by the given id",
			Action: cmd.SubnetServerList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Subnet Id",
				},
			},
		},
	}
}
