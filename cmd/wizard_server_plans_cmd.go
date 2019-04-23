package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/wizard"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpWizServerPlan prepares common resources to send request to Concerto API
func WireUpWizServerPlan(c *cli.Context) (ds *wizard.WizServerPlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = wizard.NewWizServerPlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up wizard server plan service", err)
	}

	return ds, f
}

// WizServerPlanList subcommand function
func WizServerPlanList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpWizServerPlan(c)

	checkRequiredFlags(c, []string{"app-id", "location-id", "cloud-provider-id"}, formatter)

	serverPlans, err := serverPlanSvc.GetWizServerPlanList(c.String("app-id"), c.String("location-id"), c.String("cloud-provider-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	locationsMap := LoadLocationsMapping(c)

	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		serverPlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
