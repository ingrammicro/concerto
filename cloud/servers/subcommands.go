package servers

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns servers commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists information about all the servers on this account.",
			Action: cmd.ServerList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the server identified by the given id.",
			Action: cmd.ServerShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new server.",
			Action: cmd.ServerCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the server",
				},
				cli.StringFlag{
					Name:  "ssh-profile-id",
					Usage: "Identifier of the ssh profile which the server shall use",
				},
				cli.StringFlag{
					Name:  "firewall-profile-id",
					Usage: "Identifier of the firewall profile to which the server shall use",
				},
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Identifier of the template the server shall use",
				},
				cli.StringFlag{
					Name:  "server-plan-id",
					Usage: "Identifier of the server plan in which the server shall be deployed",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which the server shall be registered",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with server",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing server",
			Action: cmd.ServerUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the server",
				},
			},
		},
		{
			Name:   "boot",
			Usage:  "Boots a server with the given id",
			Action: cmd.ServerBoot,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "reboot",
			Usage:  "Reboots a server with the given id",
			Action: cmd.ServerReboot,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "shutdown",
			Usage:  "Shuts down a server with the given id",
			Action: cmd.ServerShutdown,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "override-server",
			Usage:  "This action takes the server with the given id from a stalled state to the operational state, at the user's own risk.",
			Action: cmd.ServerOverride,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "This action decommissions the server with the given id. The server must be in a inactive, stalled or commission_stalled state.",
			Action: cmd.ServerDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "list-events",
			Usage:  "This action returns information about the events related to the server with the given id.",
			Action: cmd.EventsList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "list-operational-scripts",
			Usage:  "This action returns information about the operational scripts characterisations related to the server with the given id.",
			Action: cmd.OperationalScriptsList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "execute-script",
			Usage:  "This action initiates the execution of the script characterisation with the given id on the server with the given id.",
			Action: cmd.OperationalScriptExecute,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server-id",
					Usage: "Server Id",
				},
				cli.StringFlag{
					Name:  "script-id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "list-volumes",
			Usage:  "This action returns information about the volumes attached to the server with the given id",
			Action: cmd.ServerVolumesList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
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
					Usage: "Server Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "server",
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
					Usage: "Server Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "server",
					Hidden: true,
				},
			},
		},
	}
}
