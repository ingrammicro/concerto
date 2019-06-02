package server_arrays

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns server arrays commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists information about all the server arrays on this account",
			Action: cmd.ServerArrayList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the server array identified by the given id",
			Action: cmd.ServerArrayShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new server array",
			Action: cmd.ServerArrayCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the server",
				},
				cli.IntFlag{
					Name:  "size",
					Usage: "Number of initial servers in the server array. Value by default is 0",
				},
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Identifier of the template the server array shall use",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which the server array shall be registered",
				},
				cli.StringFlag{
					Name:  "server-plan-id",
					Usage: "Identifier of the server plan in which the server array shall be deployed",
				},
				cli.StringFlag{
					Name:  "firewall-profile-id",
					Usage: "Identifier of the firewall profile to which the server array belongs. It will take default firewall profile if it is not given",
				},
				cli.StringFlag{
					Name:  "ssh-profile-id",
					Usage: "Identifier of the ssh profile to which the server array belongs. It will take default ssh profile if it is not given",
				},
				cli.StringFlag{
					Name:  "subnet-id",
					Usage: "Identifier of the subnet to which the server array belongs. It will not be on any subnet managed by IMCO if not given",
				},
				cli.BoolFlag{
					Name:  "privateness",
					Usage: "If the server array is private, set this flag, i.e: --privateness. If it's public, do not set this flag",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with server array",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing server array",
			Action: cmd.ServerArrayUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the server array",
				},
			},
		},
		{
			Name:   "boot",
			Usage:  "This action boots all the servers in the server array with the given id. The server array must be in an inactive state",
			Action: cmd.ServerArrayBoot,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
			},
		},
		{
			Name:   "shutdown",
			Usage:  "This action shuts down all the servers in the server array identified by the given id. The server must be in a bootstrap",
			Action: cmd.ServerArrayShutdown,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
			},
		},
		{
			Name:   "empty",
			Usage:  "This action empties all servers in server array with the given id",
			Action: cmd.ServerArrayEmpty,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
			},
		},
		{
			Name:   "enlarge",
			Usage:  "This action add servers in server array with the given id",
			Action: cmd.ServerArrayEnlarge,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
				cli.IntFlag{
					Name:  "size",
					Usage: "The number of servers to add to the array, a number between 1 and 5",
				},
			},
		},
		{
			Name:   "list-servers",
			Usage:  "This action list servers in server array with the given id",
			Action: cmd.ServerArrayServerList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
				},
			},
		},
		{
			Name:   "destroy",
			Usage:  "This action destroys the server array with the given id. This action will only be allowed if the server array is empty",
			Action: cmd.ServerArrayDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Array Id",
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
					Usage: "Server Array Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "server_array",
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
					Usage: "Server Array Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "server_array",
					Hidden: true,
				},
			},
		},
	}
}
