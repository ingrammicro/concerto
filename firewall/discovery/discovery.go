package discovery

import (
	"fmt"
	"net"
)

type FirewallChain struct {
	Name   string
	Policy string
	Rules  []*FirewallRule
}

type FirewallRule struct {
	Name     string
	Target   string
	Protocol string
	Source   string
	Dports   [2]int
}

func (fc *FirewallChain) String() string {
	return fmt.Sprintf("{chain name='%s' policy='%s' rules=%v}", fc.Name, fc.Policy, fc.Rules)
}

func (fr *FirewallRule) String() string {
	return fmt.Sprintf("{target='%s' protocol='%s' source='%s' minPort=%d maxPort=%d}", fr.Target, fr.Protocol, fr.Source, fr.Dports[0], fr.Dports[1])
}

func FlattenChain(chainName string, chains []*FirewallChain, affectingRule *FirewallRule) (*FirewallChain, error) {
	var c *FirewallChain
	var chainIndex int
	for chainIndex, c = range chains {
		if c.Name == chainName {
			break
		}
	}
	if c == nil {
		return nil, fmt.Errorf("chain %s not defined or infinite recursion", chainName)
	}
	chains[chainIndex] = chains[len(chains)-1]
	chains = chains[:len(chains)-1]
	result := &FirewallChain{
		Name:   chainName,
		Policy: "DROP",
	}
	if affectingRule == nil {
		affectingRule = &FirewallRule{
			Target:   "ACCEPT",
			Protocol: "all",
			Source:   "0.0.0.0/0",
			Dports:   [2]int{1, 65535},
		}
	}
	if c.Policy == "ACCEPT" {
		result.Rules = []*FirewallRule{
			&FirewallRule{
				Target:   "ACCEPT",
				Protocol: affectingRule.Protocol,
				Source:   affectingRule.Source,
				Dports:   [2]int{affectingRule.Dports[0], affectingRule.Dports[1]},
			},
		}
		return result, nil
	}
	if c.Policy == "DROP" || c.Policy == "" {
		for _, rule := range c.Rules {
			if rule.Target == "ACCEPT" {
				r, err := intersectFirewallRules(affectingRule, rule)
				if err != nil {
					fmt.Printf("Warning: merging rules: %v", err)
				} else {
					if r != nil {
						result.Rules = append(result.Rules, r)
					}
				}
			} else if rule.Target != "DROP" {
				r, err := intersectFirewallRules(affectingRule, rule)
				if err != nil {
					fmt.Printf("Warning: merging rules: %v", err)
				} else {
					if r != nil {
						flattenedChain, err := FlattenChain(rule.Target, chains, r)
						if err != nil {
							fmt.Printf("Warning: flattening chain: %v", err)
						} else {
							result.Rules = append(result.Rules, flattenedChain.Rules...)
						}
					}
				}
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("custom policy %s as chain default policy", c.Policy)
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
			net2Size, _ := net2.Mask.Size()
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
	protocol := intersectFirewallRuleProtocol(r1.Protocol, r2.Protocol)
	if protocol == "" {
		return nil, nil
	}
	source, err := intersectFirewallRuleSource(r1.Source, r2.Source)
	if source == "" || err != nil {
		return nil, err
	}
	dports := intersectFirewallRuleDPorts(r1.Dports, r2.Dports)
	if dports[1] == 0 {
		return nil, nil
	}
	return &FirewallRule{
		Target:   "ACCEPT",
		Protocol: protocol,
		Source:   source,
		Dports:   dports,
	}, nil
}
