package brownfield

import "github.com/codegangsta/cli"

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "register",
			Usage:  "Register concerto agent within an imported brownfield Host",
			Action: cmdRegister,
		},
		{
			Name:   "configure",
			Usage:  "Configures an imported brownfield Host's FQDN, SSH access, agent services and firewall",
			Action: cmdConfigure,
		},
	}
}
