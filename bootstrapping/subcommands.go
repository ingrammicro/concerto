package bootstrapping

import (
	"github.com/codegangsta/cli"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "start",
			Usage:  "Starts a bootstrapping routine to check and execute required activities",
			Action: start,
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "time, t",
					Usage: "bootstrapping time interval (seconds)",
					Value: DefaultTimingInterval,
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops the running bootstrapping process",
			Action: stop,
		},
	}
}
