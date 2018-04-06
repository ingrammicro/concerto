// +build windows

package firewall

import (
	"fmt"

	"github.com/ingrammicro/concerto/firewall/discovery"

	"github.com/ingrammicro/concerto/utils"
)

func driverName() string {
	return "windows"
}

func apply(policy Policy) error {
	err := flush()
	if err != nil {
		return err
	}
	for i, rule := range policy.Rules {
		cidr := rule.Cidr
		if rule.Cidr == "0.0.0.0/0" {
			cidr = "any"
		}
		ruleCmd := fmt.Sprintf(
			"netsh advfirewall firewall add rule name=\"Concerto firewall %d\" dir=in action=allow remoteip=\"%s\" protocol=\"%s\" localport=\"%d-%d\"",
			i, cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
		utils.RunCmd(ruleCmd)
	}

	utils.RunCmd("netsh advfirewall set allprofiles state on")
	return nil
}

func flush() error {
	fc, err := discovery.CurrentFirewallRules()
	if err != nil {
		return err
	}
	utils.RunCmd("netsh advfirewall set allprofiles state off")
	utils.RunCmd("netsh advfirewall set allprofiles firewallpolicy allowinbound,allowoutbound")
	//utils.RunCmd("netsh advfirewall firewall delete rule name=all")
	for _, r := range fc[0].Rules {
		utils.RunCmd(fmt.Sprintf("netsh advfirewall firewall delete rule name=%q", r.Name))
	}
	return nil
}
