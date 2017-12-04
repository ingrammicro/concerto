package cmdpolling

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "register",
			Usage:  "Registers concerto agent within an imported Host",
			Action: cmdRegister,
		},
		{
			Name:   "start",
			Usage:  "Starts a polling routine to check and execute pending scripts",
			Action: cmdStart,
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "time, t",
					Usage: "Polling ping time interval (seconds)",
					Value: 30,
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops the running polling process",
			Action: cmdStop,
		},
		{
			Name:   "continuous-report-run",
			Usage:  "Runs a script and gradually report its output",
			Action: cmdContinuousReportRun,
			ArgsUsage: "script",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "lines, l",
					Usage: "Maximum lines threshold per response chunk",
					Value: utils.DefaultThresholdLines,
				},
				cli.IntFlag{
					Name:  "time, t",
					Usage: "Maximum time -seconds- threshold per response chunk",
					Value: utils.DefaultThresholdTime,
				},
				cli.IntFlag{
					Name:  "bytes, b",
					Usage: "Maximum bytes threshold per response chunk",
					Value: utils.DefaultThresholdBytes,
				},
			},
		},
	}
}
