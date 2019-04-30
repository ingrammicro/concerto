package firewall

import (
	"github.com/codegangsta/cli"
)

// SubCommands returns firewall commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "add",
			Usage:  "Adds a single firewall rule to host",
			Action: cmdAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "min-port",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "max-port",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ip-protocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "apply",
			Usage:  "Applies selected firewall rules in host",
			Action: cmdApply,
		},
		{
			Name:   "check",
			Usage:  "Checks if a firewall rule exists in host",
			Action: cmdCheck,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "min-port",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "max-port",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ip-protocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "flush",
			Usage:  "Flushes all firewall rules from host",
			Action: cmdFlush,
		},
		{
			Name:   "list",
			Usage:  "Lists all firewall rules associated to host",
			Action: cmdList,
		},
		{
			Name:   "remove",
			Usage:  "Removes a firewall rule to host",
			Action: cmdRemove,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "min-port",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "max-port",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ip-protocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates all firewall rules",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "rules",
					Usage: `JSON array in the form '[{"ip_protocol":"...", "min_port":..., "max_port":..., "cidr_ip":"..."}, ... ]'`,
				},
			},
		},
	}
}
