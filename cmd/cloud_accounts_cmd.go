package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/settings"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpCloudAccount prepares common resources to send request to Concerto API
func WireUpCloudAccount(c *cli.Context) (ds *settings.CloudAccountService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = settings.NewCloudAccountService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cloudAccount service", err)
	}

	return ds, f
}

// CloudAccountList subcommand function
func CloudAccountList(c *cli.Context) error {
	debugCmdFuncInfo(c)

	cloudAccountSvc, formatter := WireUpCloudAccount(c)
	cloudAccounts, err := cloudAccountSvc.GetCloudAccountList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudAccount data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)

	for id, ca := range cloudAccounts {
		cloudAccounts[id].CloudProviderName = cloudProvidersMap[ca.CloudProviderID]
	}

	if err = formatter.PrintList(cloudAccounts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}

	return nil
}

// CloudAccountShow subcommand function
func CloudAccountShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cloudAccount, err := cloudAccountSvc.GetCloudAccount(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudAccount data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)

	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
