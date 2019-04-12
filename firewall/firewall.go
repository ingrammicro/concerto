package firewall

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

func cmdList(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	return cmd.FirewallRuleList(c)
}

func cmdApply(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	policy := cmd.FirewallPolicyGet(c)
	// Only apply firewall if we get a non-empty set of rules
	if len(policy.Rules) > 0 {
		return Apply(*policy)
	}
	return flush()
}

func cmdFlush(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	return flush()
}

func cmdCheck(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	_, _, exists := cmd.FirewallRuleCheck(c)
	fmt.Printf("%t\n", exists)
	return nil
}

func cmdAdd(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	return cmd.FirewallRuleAdd(c)
}

func cmdUpdate(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	return cmd.FirewallRulesUpdate(c)
}

func cmdRemove(c *cli.Context) error {
	log.Debugf("Current firewall driver %s", driverName())
	return cmd.FirewallRuleRemove(c)
}
