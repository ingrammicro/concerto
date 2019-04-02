package cookbook_versions

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available cookbook versions",
			Action: cmd.CookbookVersionList,
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
	}
}
