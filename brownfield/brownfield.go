package brownfield

import "github.com/codegangsta/cli"

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "register",
			Usage:  "Register concerto agent within an imported brownfield Host",
			Action: cmdRegister,
		},
	}
}
