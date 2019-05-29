package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/storage"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpStoragePlan prepares common resources to send request to Concerto API
func WireUpStoragePlan(c *cli.Context) (ns *storage.StoragePlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = storage.NewStoragePlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up storage plan service", err)
	}

	return ns, f
}

// StoragePlanShow subcommand function
func StoragePlanShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	storagePlanSvc, formatter := WireUpStoragePlan(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	storagePlans, err := storagePlanSvc.GetStoragePlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive storage plan data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	locationsMap := LoadLocationsMapping(c)

	storagePlans.CloudProviderName = cloudProvidersMap[storagePlans.CloudProviderID]
	storagePlans.LocationName = locationsMap[storagePlans.LocationID]

	if err = formatter.PrintItem(*storagePlans); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
