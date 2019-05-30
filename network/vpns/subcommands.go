package vpns

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns VPNs commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "show",
			Usage:  "Shows information about the VPN identified by the given VPC id",
			Action: cmd.VPNShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "VPC Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new VPN for the specified VPC",
			Action: cmd.VPNCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "VPC Id",
				},
				cli.StringFlag{
					Name:  "public-ip",
					Usage: "Public Ip for the VPN",
				},
				cli.StringFlag{
					Name:  "psk",
					Usage: "Pass key of the VPN",
				},
				cli.StringFlag{
					Name:  "exposed-cidrs",
					Usage: "A list of comma separated exposed CIDRs of the VPN",
				},
				cli.StringFlag{
					Name:  "vpn-plan-id",
					Usage: "Identifier of the VPN plan",
				},
			},
		},
		{
			Name:   "destroy",
			Usage:  "Deletes a VPN of the specified VPC",
			Action: cmd.VPNDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "VPC Id",
				},
			},
		},
		{
			Name:   "list-plans",
			Usage:  "Lists VPN plans of the specified VPC",
			Action: cmd.VPNPlanList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "VPC Id",
				},
			},
		},
	}
}
