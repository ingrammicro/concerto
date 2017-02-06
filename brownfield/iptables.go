package brownfield

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	fw "github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

func apply(p *fw.Policy) error {
	utils.RunCmd("/sbin/iptables -w -F INPUT")
	return p.Apply()
}

func obtainCurrentFirewallRules(f format.Formatter) []*FirewallChain {
	output, err := exec.Command("/sbin/iptables", "-L", "-n", "-v").Output()
	if err != nil {
		f.PrintFatal("Error happened running iptables list command", err)
	}
	chains, err := ParseIptablesOutput(string(output))
	if err != nil {
		f.PrintFatal("Error happened parsing iptables configuration", err)
	}
	return chains
}

func ParseIptablesOutput(output string) ([]*FirewallChain, error) {
	var chains []*FirewallChain
	cs := strings.Split(output, "\n\n")
	fmt.Printf("Found %d chains\n", len(cs))
	for _, c := range cs {
		chain, err := ParseIptablesChain(c)
		if err == nil {
			chains = append(chains, chain)
		} else {
			fmt.Printf("Warning: error occurred while parsing iptables chain: %v\n", err)
		}
	}
	return chains, nil
}

var iptablesChainHeaderRegexp = regexp.MustCompile(`\AChain (?P<name>[a-zA-Z0-9-]+) \((policy (?P<policy>[a-zA-Z0-9-]+) )?`)

func ParseIptablesChain(c string) (*FirewallChain, error) {
	lines := strings.Split(c, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("chain output has no header")
	}
	if len(lines) < 2 {
		return nil, fmt.Errorf("chain output has no rules header")
	}
	header := lines[0]
	chain := &FirewallChain{}
	matched, err := regexp.Match(iptablesChainHeaderRegexp.String(), []byte(header))
	if err != nil {
		return nil, fmt.Errorf("cannot parse chain header '%s': %v", header, err)
	}
	if !matched {
		return nil, fmt.Errorf("cannot parse chain header '%s'", header)
	}
	match := iptablesChainHeaderRegexp.FindStringSubmatch(header)
	fmt.Sprintf("%v\n", match)
	for i, name := range iptablesChainHeaderRegexp.SubexpNames() {
		switch name {
		case "name":
			chain.name = match[i]
		case "policy":
			chain.policy = match[i]
		}
	}
	if chain.name == "" {
		return nil, fmt.Errorf("found no name for chain in '%s'", header)
	}
	rules := lines[2:]
	for _, r := range rules {
		if r != "" {
			rule, err := ParseIptablesRule(r)
			if err != nil {
				fmt.Printf("Warning: cannot parse rule for chain %s : %v\n", chain.name, err)
			} else {
				if rule != nil {
					chain.rules = append(chain.rules, rule)
				}
			}
		}
	}
	return chain, nil
}

var iptablesRuleFieldSeparator = regexp.MustCompile("[[:blank:]]+")
var iptablesRuleDPortRegexp = regexp.MustCompile(`(tcp|udp) dpts?:(?P<minPort>\d+)(:(?P<maxPort>\d+))?`)
var iptablesRuleStateRegexp = regexp.MustCompile(`state [[:alpha:]]+(,[[:alpha:]]+)*`)
var iptablesRuleStatsInfoRegexp = regexp.MustCompile(`^ ?\d+[A-Z]? \d+[A-Z]? `)

func ParseIptablesRule(r string) (*FirewallRule, error) {
	r = iptablesRuleFieldSeparator.ReplaceAllLiteralString(r, " ")
	r = iptablesRuleStatsInfoRegexp.ReplaceAllLiteralString(r, "")
	fields := iptablesRuleFieldSeparator.Split(r, 8)
	if len(fields) < 8 {
		return nil, fmt.Errorf("rule '%s' has too few fields", r)
	}
	matchString := iptablesRuleStateRegexp.FindString(r)
	if len(matchString) > 0 { //state condition, we ignore it
		return nil, nil
	}
	if fields[3] == "lo" { // incoming interface is localhost
		return nil, nil
	}
	var dports [2]int
	matchString = iptablesRuleDPortRegexp.FindString(r)
	if len(matchString) == 0 {
		dports = [2]int{1, 65535}
	} else {
		match := iptablesRuleDPortRegexp.FindStringSubmatch(r)
		dports = [2]int{0, 0}
		for i, name := range iptablesRuleDPortRegexp.SubexpNames() {
			var err error
			switch name {
			case "minPort":
				dports[0], err = strconv.Atoi(match[i])
				if err != nil {
					return nil, fmt.Errorf("rule '%s' has invalid destination %s specification '%s'", r, name, match[i])
				}
			case "maxPort":
				dports[1], _ = strconv.Atoi(match[i])
				if err != nil {
					dports[1] = 0
				}
			}
		}
		if dports[1] == 0 {
			dports[1] = dports[0]
		}
	}
	rule := &FirewallRule{
		target:   fields[0],
		protocol: fields[1],
		source:   fields[5],
		dports:   dports,
	}
	return rule, nil
}
