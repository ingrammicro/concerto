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
					Name:  "interval, i",
					Usage: "The frequency (in seconds) at which the bootstrapping runs",
					Value: DefaultTimingInterval,
				},
				cli.Int64Flag{
					Name:  "splay, s",
					Usage: "A random number between zero and splay that is added to interval (seconds)",
					Value: DefaultTimingSplay,
				},
				cli.IntFlag{
					Name:  "lines, l",
					Usage: "Maximum lines threshold per response chunk",
					Value: DefaultThresholdLines,
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
