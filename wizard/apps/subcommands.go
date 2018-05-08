package apps

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands provides operations about apps
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Apps.",
			Action: cmd.AppList,
		},
		{
			Name:   "deploy",
			Usage:  "Deploys the App with the given id as a server on the cloud.",
			Action: cmd.AppDeploy,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier of the App which will be deployed",
				},
				cli.StringFlag{
					Name:  "location_id",
					Usage: "Identifier of the Location on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "cloud_account_id",
					Usage: "Identifier of the Cloud Account with which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "server_plan_id",
					Usage: "Identifier of the Server Plan on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "hostname",
					Usage: "A hostname for the cloud server to deploy",
				},
			},
		},
	}
}
