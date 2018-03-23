package brownfield

import (
	"encoding/json"
	"fmt"
	"strings"

	fw "github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/firewall/discovery"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

func configureConcertoFirewall(cs *utils.HTTPConcertoservice, f format.Formatter) {
	chains, err := discovery.CurrentFirewallRules()
	if err != nil {
		f.PrintFatal("Cannot obtain current firewall rules", err)
	}
	flattenedInputChain, err := discovery.FlattenChain("INPUT", chains, nil)
	if err != nil {
		f.PrintFatal("Cannot flatten firewall INPUT chain", err)
	}
	fmt.Printf("After flattening chain: %d rules\n", len(flattenedInputChain.Rules))
	policy, err := startFirewallMapping(cs, flattenedInputChain.Rules)
	if err != nil {
		f.PrintFatal("Error starting the firewall mapping", err)
	}
	err = apply(policy)
	if err != nil {
		f.PrintFatal("Applying Concerto firewall", err)
	}
}

func startFirewallMapping(cs *utils.HTTPConcertoservice, rules []*discovery.FirewallRule) (p *fw.Policy, err error) {
	payload := convertFirewallChainToPayload(rules)
	fmt.Printf("DEBUG: Sending following firewall profile: %+v\n", payload)
	body, status, err := cs.Post("/cloud/firewall_profile", &payload)
	if err != nil {
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %s", status, string(body))
		return
	}
	responseData := &fw.FirewallProfile{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}
	p = &(responseData.Profile)
	return
}

func convertFirewallChainToPayload(rules []*discovery.FirewallRule) map[string]interface{} {
	fpRules := []interface{}{}
	for _, r := range rules {
		newRules := convertRuleToPayload(r)
		if newRules != nil {
			fpRules = append(fpRules, newRules...)
		}
	}
	fp := map[string]interface{}{
		"firewall_profile": map[string]interface{}{
			"rules": fpRules,
		},
	}
	return fp
}

func convertRuleToPayload(rule *discovery.FirewallRule) []interface{} {
	var rules []interface{}
	protocol := strings.ToLower(rule.Protocol)
	if protocol != "all" && protocol != "tcp" && protocol != "udp" {
		return nil
	}
	if protocol == "all" || protocol == "tcp" {
		rules = append(rules,
			fw.Rule{
				Name:     rule.Name,
				Protocol: "tcp",
				Cidr:     rule.Source,
				MinPort:  rule.Dports[0],
				MaxPort:  rule.Dports[1],
			})
	}
	if protocol == "all" || protocol == "udp" {
		rules = append(rules,
			fw.Rule{
				Name:     rule.Name,
				Protocol: "udp",
				Cidr:     rule.Source,
				MinPort:  rule.Dports[0],
				MaxPort:  rule.Dports[1],
			})
	}
	return rules
}
