package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/dispatcher"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpDispatcher prepares common resources to send request to API
func WireUpDispatcher(c *cli.Context) (ds *dispatcher.DispatcherService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = dispatcher.NewDispatcherService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up dispatcher service", err)
	}

	return ds, f
}
