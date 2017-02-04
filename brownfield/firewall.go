package brownfield

import (
	"encoding/json"
	"fmt"
	"net"

	fw "github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

type FirewallChain struct {
	name   string
	policy string
	rules  []*FirewallRule
}

type FirewallRule struct {
	target   string
	protocol string
	source   string
	dports   [2]int
}

func configureConcertoFirewall(cs *utils.HTTPConcertoservice, f format.Formatter) {
	chains := obtainCurrentFirewallRules(f)
	fmt.Printf("Parsed chains: %v\n", chains)
	flattenedInputChain, err := buildFlattenedChain("INPUT", chains, nil)
	if err != nil {
		f.PrintFatal("Cannot flatten firewall INPUT chain", err)
	}
	fmt.Printf("Flattened chain: %v\n", flattenedInputChain)
	policy, err := startFirewallMapping(cs, flattenedInputChain.rules)
	if err != nil {
		f.PrintFatal("Error starting the firewall mapping", err)
	}
	fmt.Printf("Policy to apply: %v\n", policy)
}

func (fc *FirewallChain) String() string {
	return fmt.Sprintf("{chain name='%s' policy='%s' rules=%v}", fc.name, fc.policy, fc.rules)
}

func (fr *FirewallRule) String() string {
	return fmt.Sprintf("{target='%s' protocol='%s' source='%s' minPort=%d maxPort=%d}", fr.target, fr.protocol, fr.source, fr.dports[0], fr.dports[1])
}

func buildFlattenedChain(chainName string, chains []*FirewallChain, affectingRule *FirewallRule) (*FirewallChain, error) {
	var c *FirewallChain
	var chainIndex int
	for chainIndex, c = range chains {
		if c.name == chainName {
			break
		}
	}
	if c == nil {
		return nil, fmt.Errorf("chain %s not defined or infinite recursion", chainName)
	}
	chains[chainIndex] = chains[len(chains)-1]
	chains = chains[:len(chains)-1]
	result := &FirewallChain{
		name:   chainName,
		policy: "DROP",
	}
	if affectingRule == nil {
		affectingRule = &FirewallRule{
			target:   "ACCEPT",
			protocol: "all",
			source:   "0.0.0.0/0",
			dports:   [2]int{1, 65535},
		}
	}
	if c.policy == "ACCEPT" {
		result.rules = []*FirewallRule{
			&FirewallRule{
				target:   "ACCEPT",
				protocol: affectingRule.protocol,
				source:   affectingRule.source,
				dports:   [2]int{affectingRule.dports[0], affectingRule.dports[1]},
			},
		}
		return result, nil
	}
	if c.policy == "DROP" || c.policy == "" {
		for _, rule := range c.rules {
			if rule.target == "ACCEPT" {
				r, err := intersectFirewallRules(affectingRule, rule)
				if err != nil {
					fmt.Printf("Warning: merging rules: %v", err)
				} else {
					if r != nil {
						result.rules = append(result.rules, r)
					}
				}
			} else if rule.target != "DROP" {
				r, err := intersectFirewallRules(affectingRule, rule)
				if err != nil {
					fmt.Printf("Warning: merging rules: %v", err)
				} else {
					if r != nil {
						flattenedChain, err := buildFlattenedChain(rule.target, chains, r)
						if err != nil {
							fmt.Printf("Warning: flattening chain: %v", err)
						} else {
							result.rules = append(result.rules, flattenedChain.rules...)
						}
					}
				}
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("custom policy %s as chain default policy", c.policy)
}

func intersectFirewallRuleProtocol(p1, p2 string) string {
	if p1 == p2 {
		return p1
	}
	if p1 == "all" {
		return p2
	}
	if p2 == "all" {
		return p1
	}
	return ""
}

func intersectFirewallRuleSource(s1, s2 string) (string, error) {
	if s1 == s2 {
		return s1, nil
	}
	ip1, net1, err := net.ParseCIDR(s1)
	if err != nil {
		return "", fmt.Errorf("invalid source: %v", err)
	}
	ip1 = ip1.To4()
	if ip1 == nil {
		return "", fmt.Errorf("invalid source: %v is not an IPv4", s1)
	}
	ip2, net2, err := net.ParseCIDR(s2)
	if err != nil {
		return "", fmt.Errorf("invalid source: %v", err)
	}
	ip2 = ip2.To4()
	if ip2 == nil {
		return "", fmt.Errorf("invalid source: %v is not an IPv4", s2)
	}
	if net1.Contains(ip2) {
		if net2.Contains(ip1) {
			net1Size, _ := net1.Mask.Size()
			net2Size, _ := net1.Mask.Size()
			if net1Size > net2Size {
				return s1, nil
			}
			return s2, nil
		}
		return s2, nil
	}
	if net2.Contains(ip1) {
		return s1, nil
	}
	return "", nil
}

func intersectFirewallRuleDPorts(dports1, dports2 [2]int) (dports [2]int) {
	if dports1[0] > dports2[0] {
		dports[0] = dports1[0]
	} else {
		dports[0] = dports2[0]
	}
	if dports2[1] < dports1[1] {
		dports[1] = dports2[1]
	} else {
		dports[1] = dports1[1]
	}
	if dports[0] > dports[1] {
		dports[0] = 0
		dports[1] = 0
	}
	return
}

func intersectFirewallRules(r1, r2 *FirewallRule) (*FirewallRule, error) {
	protocol := intersectFirewallRuleProtocol(r1.protocol, r2.protocol)
	if protocol == "" {
		return nil, nil
	}
	source, err := intersectFirewallRuleSource(r1.source, r2.source)
	if source == "" || err != nil {
		return nil, err
	}
	dports := intersectFirewallRuleDPorts(r1.dports, r2.dports)
	if dports[1] == 0 {
		return nil, nil
	}
	return &FirewallRule{
		target:   "ACCEPT",
		protocol: protocol,
		source:   source,
		dports:   dports,
	}, nil
}

func startFirewallMapping(cs *utils.HTTPConcertoservice, rules []*FirewallRule) (p *fw.Policy, err error) {
	payload := convertFirewallChainToPayload(rules)
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

func convertFirewallChainToPayload(rules []*FirewallRule) map[string]interface{} {
	fpRules := []interface{}{}
	for _, r := range rules {
		newRules := convertRuleToPayload(r)
		fpRules = append(fpRules, newRules...)
	}
	fp := map[string]interface{}{
		"firewall_profile": map[string]interface{}{
			"rules": fpRules,
		},
	}
	return fp
}

func convertRuleToPayload(rule *FirewallRule) (rules []interface{}) {
	protocol := rule.protocol
	if protocol != "all" && protocol != "tcp" && protocol != "udp" {
		return
	}
	if protocol == "all" || protocol == "tcp" {
		rules = append(rules,
			fw.Rule{
				Protocol: "tcp",
				Cidr:     rule.source,
				MinPort:  rule.dports[0],
				MaxPort:  rule.dports[1],
			})
	}
	if protocol == "all" || protocol == "udp" {
		rules = append(rules,
			fw.Rule{
				Protocol: "udp",
				Cidr:     rule.source,
				MinPort:  rule.dports[0],
				MaxPort:  rule.dports[1],
			})
	}
	return
}
