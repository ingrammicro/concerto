package firewall_profiles

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns firewall profiles commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing firewall profiles",
			Action: cmd.FirewallProfileList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the firewall profile identified by the given id.",
			Action: cmd.FirewallProfileShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a a firewall profile with the given parameters.",
			Action: cmd.FirewallProfileCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the firewall profile",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the firewall profile",
				},
				cli.StringFlag{
					Name:  "rules",
					Usage: "Set of rules of the firewall profile",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with firewall profile",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates the firewall profile identified by the given id with the given parameters.",
			Action: cmd.FirewallProfileUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the firewall profile",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the firewall profile",
				},
				cli.StringFlag{
					Name:  "rules",
					Usage: "Set of rules of the firewall profile",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Destroy a firewall profile",
			Action: cmd.FirewallProfileDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
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
					Usage: "Firewall profile Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "firewall_profile",
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
					Usage: "Firewall profile Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "firewall_profile",
					Hidden: true,
				},
			},
		},
	}
}
