package main

import (
	"fmt"
	"os"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/audit"
	"github.com/ingrammicro/concerto/blueprint/scripts"
	"github.com/ingrammicro/concerto/blueprint/services"
	"github.com/ingrammicro/concerto/blueprint/templates"
	"github.com/ingrammicro/concerto/brownfield"
	cl_prov "github.com/ingrammicro/concerto/cloud/cloud_providers"
	"github.com/ingrammicro/concerto/cloud/generic_images"
	"github.com/ingrammicro/concerto/cloud/saas_providers"
	"github.com/ingrammicro/concerto/cloud/server_plan"
	"github.com/ingrammicro/concerto/cloud/servers"
	"github.com/ingrammicro/concerto/cloud/ssh_profiles"
	"github.com/ingrammicro/concerto/cmdpolling"
	"github.com/ingrammicro/concerto/converge"
	"github.com/ingrammicro/concerto/dispatcher"
	"github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/network/firewall_profiles"
	"github.com/ingrammicro/concerto/settings/cloud_accounts"
	"github.com/ingrammicro/concerto/setup"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
	"github.com/ingrammicro/concerto/wizard/apps"
	"github.com/ingrammicro/concerto/wizard/cloud_providers"
	"github.com/ingrammicro/concerto/wizard/locations"
	"github.com/ingrammicro/concerto/wizard/server_plans"
)

// ServerCommands stores Commands array for Server mode
var ServerCommands = []cli.Command{
	{
		Name:  "firewall",
		Usage: "Manages Firewall Policies within a Host",
		Subcommands: append(
			firewall.SubCommands(),
		),
	},
	{
		Name:  "scripts",
		Usage: "Manages Execution Scripts within a Host",
		Subcommands: append(
			dispatcher.SubCommands(),
		),
	},
	{
		Name:   "converge",
		Usage:  "Converges Host to original Blueprint",
		Action: converge.CmbConverge,
	},
	{
		Name:  "brownfield",
		Usage: "Manages registration and configuration within an imported brownfield Host",
		Subcommands: append(
			brownfield.SubCommands(),
		),
	},
	{
		Name:  "polling",
		Usage: "Manages polling commands",
		Subcommands: append(
			cmdpolling.SubCommands(),
		),
	},
}

// BlueprintCommands stores Commands array for Blueprint functionalities
var BlueprintCommands = []cli.Command{
	{
		Name:  "scripts",
		Usage: "Allow the user to manage the scripts they want to run on the servers",
		Subcommands: append(
			scripts.SubCommands(),
		),
	},
	{
		Name:  "services",
		Usage: "Provides information on services",
		Subcommands: append(
			services.SubCommands(),
		),
	},
	{
		Name:  "templates",
		Usage: "Provides information on templates",
		Subcommands: append(
			templates.SubCommands(),
		),
	},
}

// CloudCommands stores Commands array for Cloud functionalities
var CloudCommands = []cli.Command{
	{
		Name:  "servers",
		Usage: "Provides information on servers",
		Subcommands: append(
			servers.SubCommands(),
		),
	},
	{
		Name:  "generic_images",
		Usage: "Provides information on generic images",
		Subcommands: append(
			generic_images.SubCommands(),
		),
	},
	{
		Name:  "ssh_profiles",
		Usage: "Provides information on SSH profiles",
		Subcommands: append(
			ssh_profiles.SubCommands(),
		),
	},
	{
		Name:  "cloud_providers",
		Usage: "Provides information on cloud providers",
		Subcommands: append(
			cl_prov.SubCommands(),
		),
	},
	{
		Name:  "server_plans",
		Usage: "Provides information on server plans",
		Subcommands: append(
			server_plan.SubCommands(),
		),
	},
	{
		Name:  "saas_providers",
		Usage: "Provides information about SAAS providers",
		Subcommands: append(
			saas_providers.SubCommands(),
		),
	},
}

// NetCommands stores Commands array for Network functionalities
var NetCommands = []cli.Command{
	{
		Name:  "firewall_profiles",
		Usage: "Provides information about firewall profiles",
		Subcommands: append(
			firewall_profiles.SubCommands(),
		),
	},
}

// SettingsCommands stores Commands array for Settings functionalities
var SettingsCommands = []cli.Command{
	{
		Name:  "cloud_accounts",
		Usage: "Provides information about cloud accounts",
		Subcommands: append(
			cloud_accounts.SubCommands(),
		),
	},
}

// WizardCommands stores Commands array for Wizard functionalities
var WizardCommands = []cli.Command{
	{
		Name:  "apps",
		Usage: "Provides information about apps",
		Subcommands: append(
			apps.SubCommands(),
		),
	},
	{
		Name:  "cloud_providers",
		Usage: "Provides information about cloud providers",
		Subcommands: append(
			cloud_providers.SubCommands(),
		),
	},
	{
		Name:  "locations",
		Usage: "Provides information about locations",
		Subcommands: append(
			locations.SubCommands(),
		),
	},
	{
		Name:  "server_plans",
		Usage: "Provides information about server plans",
		Subcommands: append(
			server_plans.SubCommands(),
		),
	},
}

// ClientCommands stores Commands array for Client mode
var ClientCommands = []cli.Command{
	{
		Name:      "setup",
		ShortName: "se",
		Usage:     "Configures and setups concerto cli environment",
		Subcommands: append(
			setup.SubCommands(),
		),
	},
	{
		Name:      "events",
		ShortName: "ev",
		Usage:     "Events allow the user to track their actions and the state of their servers",
		Subcommands: append(
			audit.SubCommands(),
		),
	},

	{
		Name:      "blueprint",
		ShortName: "bl",
		Usage:     "Manages blueprint commands for scripts, services and templates",
		Subcommands: append(
			BlueprintCommands,
		),
	},
	{
		Name:      "cloud",
		ShortName: "clo",
		Usage:     "Manages cloud related commands for servers, generic images, ssh profiles, cloud providers, server plans and Saas providers",
		Subcommands: append(
			CloudCommands,
		),
	},
	{
		Name:      "network",
		ShortName: "net",
		Usage:     "Manages network related commands for firewall profiles",
		Subcommands: append(
			NetCommands,
		),
	},
	{
		Name:      "settings",
		ShortName: "set",
		Usage:     "Provides settings for cloud accounts",
		Subcommands: append(
			SettingsCommands,
		),
	},
	{
		Name:      "wizard",
		ShortName: "wiz",
		Usage:     "Manages wizard related commands for apps, locations, cloud providers, server plans",
		Subcommands: append(
			WizardCommands,
		),
	},
}

var appFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "debug, D",
		Usage: "Enable debug mode",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CA_CERT",
		Name:   "ca-cert",
		Usage:  "CA to verify remote connections",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CLIENT_CERT",
		Name:   "client-cert",
		Usage:  "Client cert to use for Concerto",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CLIENT_KEY",
		Name:   "client-key",
		Usage:  "Private key used in client Concerto auth",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CONFIG",
		Name:   "concerto-config",
		Usage:  "Concerto Config File",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_ENDPOINT",
		Name:   "concerto-endpoint",
		Usage:  "Concerto Endpoint",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_URL",
		Name:   "concerto-url",
		Usage:  "Concerto Web URL",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_BROWNFIELD_TOKEN",
		Name:   "concerto-brownfield-token",
		Usage:  "Concerto Brownfield Token",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_COMMAND_POLLING_TOKEN",
		Name:   "concerto-command-polling-token",
		Usage:  "Concerto Command Polling Token",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_SERVER_ID",
		Name:   "concerto-server-id",
		Usage:  "Concerto Server ID",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_FORMATTER",
		Name:   "formatter",
		Usage:  "Output formatter [ text | json ] ",
		Value:  "text",
	},
}

func excludeFlags(visibleFlags []cli.Flag, arr []string) (flags []cli.Flag) {
	for _, flag := range visibleFlags {
		bFound := false
		for _, a := range arr {
			if a == flag.GetName() {
				bFound = true
				break
			}
		}
		if !bFound {
			flags = append(flags, flag)
		}
	}
	return
}

func cmdNotFound(c *cli.Context, command string) {
	log.Fatalf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
}

func prepareFlags(c *cli.Context) error {

	if c.Bool("debug") {
		err := os.Setenv("DEBUG", "1")
		if err != nil {
			log.Errorf("Error setting debug mode: %s", err)
			return err
		}
		log.SetOutput(os.Stderr)
		log.SetLevel(log.DebugLevel)
	}

	// try to read configuration
	config, err := utils.InitializeConcertoConfig(c)
	if err != nil {
		log.Errorf("Error reading Concerto configuration: %s", err)
		return err
	}

	// validate formatter
	if c.String("formatter") != "text" && c.String("formatter") != "json" {
		log.Errorf("Unrecognized formatter %s. Please, use one of [ text | json ]", c.String("formatter"))
		return fmt.Errorf("Unrecognized formatter %s. Please, use one of [ text | json ]", c.String("formatter"))
	}
	format.InitializeFormatter(c.String("formatter"), os.Stdout)

	if config.IsAgentMode() {
		log.Debug("Setting server commands to concerto")
		c.App.Commands = ServerCommands
	} else {
		log.Debug("Setting client commands to concerto")
		c.App.Commands = ClientCommands

		// Excluding Server/Agent contextual flags
		c.App.Flags = excludeFlags(c.App.VisibleFlags(), []string{"concerto-brownfield-token", "concerto-command-polling-token", "concerto-server-id"})
	}

	sort.Sort(cli.CommandsByName(c.App.Commands))
	sort.Sort(cli.FlagsByName(c.App.Flags))

	// hack: substitute commands in category ... we should evaluate cobra/viper
	cat := c.App.Categories()

	for _, category := range cat {
		category.Commands = category.Commands[:0]
	}

	for _, command := range c.App.Commands {
		cat = cat.AddCommand(command.Category, command)
	}
	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "concerto"
	app.Author = "Concerto Contributors"
	app.Email = "https://github.com/ingrammicro/concerto"

	app.CommandNotFound = cmdNotFound
	app.Usage = "Manages communication between Host and IMCO Platform"
	app.Version = utils.VERSION

	// set client commands by default to populate categories
	app.Commands = ClientCommands

	app.Flags = appFlags

	app.Before = prepareFlags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
