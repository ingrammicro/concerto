package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/polling"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpPolling prepares common resources to send request to Concerto API
func WireUpPolling(c *cli.Context) (ps *polling.PollingService) {

	formatter := format.GetFormatter()
	config, err := utils.GetConcertoConfig()
	if err != nil {
		formatter.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		formatter.PrintFatal("Couldn't wire up concerto service", err)
	}
	ps, err = polling.NewPollingService(hcs)
	if err != nil {
		formatter.PrintFatal("Couldn't wire up polling service", err)
	}

	return ps
}
