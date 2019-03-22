package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpBootstrapping prepares common resources to send request to API
func WireUpBootstrapping(c *cli.Context) (ds *blueprint.BootstrappingService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = blueprint.NewBootstrappingService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up serverPlan service", err)
	}

	return ds, f
}