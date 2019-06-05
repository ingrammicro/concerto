package scripts

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns scripts commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available scripts",
			Action: cmd.ScriptsList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific script",
			Action: cmd.ScriptShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new script to be used in the templates",
			Action: cmd.ScriptCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the script",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the script's purpose",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "The script's code",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The names of the script's parameters",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with script",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing script",
			Action: cmd.ScriptUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the script",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the script's purpose",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "The script's code",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The names of the script's parameters",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a script",
			Action: cmd.ScriptDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "add-attachment",
			Usage:  "Adds an attachment to a script",
			Action: cmd.ScriptAttachmentAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the attachment",
				},
				cli.StringFlag{
					Name:  "filepath",
					Usage: "path to attachment file",
				},
			},
		},
		{
			Name:   "list-attachments",
			Usage:  "List the attachments a script has",
			Action: cmd.ScriptAttachmentList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
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
					Usage: "Script Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "script",
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
					Usage: "Script Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "script",
					Hidden: true,
				},
			},
		},
	}
}
