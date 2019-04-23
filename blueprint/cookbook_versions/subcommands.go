package cookbook_versions

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns cookbook versions commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available cookbook versions",
			Action: cmd.CookbookVersionList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific cookbook version",
			Action: cmd.CookbookVersionShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cookbook version Id",
				},
			},
		},
		{
			Name:   "upload",
			Usage:  "Uploads a new cookbook version",
			Action: cmd.CookbookVersionUpload,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "filepath",
					Usage: "path to cookbook version file",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with cookbook version",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a cookbook version",
			Action: cmd.CookbookVersionDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cookbook version Id",
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
					Usage: "Cookbook version Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "cookbook_version",
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
					Usage: "Cookbook version Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "cookbook_version",
					Hidden: true,
				},
			},
		},
	}
}
