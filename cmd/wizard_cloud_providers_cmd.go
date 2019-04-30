package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/wizard"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpWizCloudProvider prepares common resources to send request to Concerto API
func WireUpWizCloudProvider(c *cli.Context) (cs *wizard.WizCloudProvidersService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = wizard.NewWizCloudProvidersService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cloudProvider service", err)
	}

	return cs, f
}

// WizCloudProviderList subcommand function
func WizCloudProviderList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudProviderSvc, formatter := WireUpWizCloudProvider(c)

	checkRequiredFlags(c, []string{"app-id", "location-id"}, formatter)

	cloudProviders, err := cloudProviderSvc.GetWizCloudProviderList(c.String("app-id"), c.String("location-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	if err = formatter.PrintList(cloudProviders); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
