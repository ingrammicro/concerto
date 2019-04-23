package main

import (
	"fmt"
	"github.com/ingrammicro/concerto/blueprint"
	"github.com/ingrammicro/concerto/bootstrapping"
	"github.com/ingrammicro/concerto/brownfield"
	"github.com/ingrammicro/concerto/cloud"
	"github.com/ingrammicro/concerto/cmdpolling"
	"github.com/ingrammicro/concerto/converge"
	"github.com/ingrammicro/concerto/labels"
	"github.com/ingrammicro/concerto/network"
	"github.com/ingrammicro/concerto/settings"
	"github.com/ingrammicro/concerto/wizard"
	"os"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/audit"
	"github.com/ingrammicro/concerto/dispatcher"
	"github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

var serverCommands = []cli.Command{
	{
		Name:        "bootstrap",
		Usage:       "Manages bootstrapping commands",
		Subcommands: append(bootstrapping.SubCommands()),
	},
	{
		Name:        "brownfield",
		Usage:       "Manages registration and configuration within an imported brownfield Host",
		Subcommands: append(brownfield.SubCommands()),
	},
	{
		Name:   "converge",
		Usage:  "Converges Host to original Blueprint",
		Action: converge.CmbConverge,
	},
	{
		Name:        "firewall",
		Usage:       "Manages Firewall Policies within a Host",
		Subcommands: append(firewall.SubCommands()),
	},
	{
		Name:        "polling",
		Usage:       "Manages polling commands",
		Subcommands: append(cmdpolling.SubCommands()),
	},
	{
		Name:        "scripts",
		Usage:       "Manages Execution Scripts within a Host",
		Subcommands: append(dispatcher.SubCommands()),
	},
}

var clientCommands = []cli.Command{
	{
		Name:        "blueprint",
		ShortName:   "bl",
		Usage:       "Manages blueprint commands for scripts, services and templates",
		Subcommands: append(blueprint.SubCommands()),
	},
	{
		Name:        "cloud",
		ShortName:   "clo",
		Usage:       "Manages cloud related commands for servers, generic images, ssh profiles, cloud providers, server plans and Saas providers",
		Subcommands: append(cloud.SubCommands()),
	},
	{
		Name:        "events",
		ShortName:   "ev",
		Usage:       "Events allow the user to track their actions and the state of their servers",
		Subcommands: append(audit.SubCommands()),
	},
	{
		Name:        "labels",
		ShortName:   "lbl",
		Usage:       "Provides information about labels",
		Subcommands: append(labels.SubCommands()),
	},
	{
		Name:        "network",
		ShortName:   "net",
		Usage:       "Manages network related commands for firewall profiles",
		Subcommands: append(network.SubCommands()),
	},
	{
		Name:        "settings",
		ShortName:   "set",
		Usage:       "Provides settings for cloud accounts",
		Subcommands: append(settings.SubCommands()),
	},
	{
		Name:        "wizard",
		ShortName:   "wiz",
		Usage:       "Manages wizard related commands for apps, locations, cloud providers, server plans",
		Subcommands: append(wizard.SubCommands()),
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
		if err := os.Setenv("DEBUG", "1"); err != nil {
			log.Errorf("Couldn't set environment debug mode: %s", err)
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
		return fmt.Errorf("unrecognized formatter %s. Please, use one of [ text | json ]", c.String("formatter"))
	}
	format.InitializeFormatter(c.String("formatter"), os.Stdout)

	if config.IsAgentMode() {
		log.Debug("Setting server commands to concerto")
		c.App.Commands = serverCommands
	} else {
		log.Debug("Setting client commands to concerto")
		c.App.Commands = clientCommands

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
	app.Commands = clientCommands

	app.Flags = appFlags

	app.Before = prepareFlags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
