package cmdpolling

import (
	"github.com/codegangsta/cli"
)

// SubCommands return Polling subcommands
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
					Name:  "longTime, l",
					Usage: "Polling ping long time interval (seconds)",
					Value: DefaultPollingPingTimingIntervalLong,
				},
				cli.Int64Flag{
					Name:  "shortTime, s",
					Usage: "Polling ping short time interval (seconds)",
					Value: DefaultPollingPingTimingIntervalShort,
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops the running polling process",
			Action: cmdStop,
		},
		{
			Name:      "continuous-report-run",
			Usage:     "Runs a script and gradually report its output",
			Action:    cmdContinuousReportRun,
			ArgsUsage: "script",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "time, t",
					Usage: "Maximum time -seconds- threshold per response chunk",
					Value: DefaultThresholdTime,
				},
			},
		},
	}
}
