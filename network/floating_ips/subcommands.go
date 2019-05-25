package floating_ips

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns floating IPs commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing floating IPs",
			Action: cmd.FloatingIPList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server-id",
					Usage: "Identifier of a server to return only the floating IPs that are attached with that server",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the floating IP identified by the given id",
			Action: cmd.FloatingIPShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Floating IP Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new floating IP",
			Action: cmd.FloatingIPCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the floating IP",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which the floating IP is",
				},
				cli.StringFlag{
					Name:  "realm-id",
					Usage: "Identifier of the realm in which the floating IP is",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with floating IP",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing floating IP identified by the given id",
			Action: cmd.FloatingIPUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Floating IP Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the floating IP",
				},
			},
		},
		{
			Name:   "attach",
			Usage:  "Attaches the floating IP to server",
			Action: cmd.FloatingIPAttach,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Floating IP Id",
				},
				cli.StringFlag{
					Name:  "server-id",
					Usage: "Identifier of the server to attach the floating IP",
				},
			},
		},
		{
			Name:   "detach",
			Usage:  "Detaches a floating IP from server",
			Action: cmd.FloatingIPDetach,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Floating IP Id",
				},
			},
		},
		{
			Name:   "destroy",
			Usage:  "Deletes a floating IP",
			Action: cmd.FloatingIPDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Floating IP Id",
				},
			},
		},
		{
			Name:   "discard",
			Usage:  "Discards a floating IP but does not delete it from the cloud provider",
			Action: cmd.FloatingIPDiscard,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Floating IP Id",
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
					Usage: "Floating IP Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "floating_ip",
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
					Usage: "Floating IP Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "floating_ip",
					Hidden: true,
				},
			},
		},
	}
}
