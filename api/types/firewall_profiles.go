package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type FirewallProfile struct {
	ID           string `json:"id" header:"ID"`
	Name         string `json:"name,omitempty" header:"NAME"`
	Description  string `json:"description,omitempty" header:"DESCRIPTION"`
	Default      bool   `json:"default,omitempty" header:"DEFAULT"`
	Rules        []Rule `json:"rules,omitempty" header:"RULES" show:"nolist"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}

type Rule struct {
	Protocol string `json:"ip_protocol" header:"IP_PROTOCOL"`
	MinPort  int    `json:"min_port" header:"MIN_PORT"`
	MaxPort  int    `json:"max_port" header:"MAX_PORT"`
	CidrIP   string `json:"source" header:"SOURCE"`
}

var firewallProfileRulesRegexp = regexp.MustCompile(`(?P<ip_protocol>\w{3})\/(?P<min_port>\d+)(?:-(?P<max_port>\d+)?)?:(?P<source>[a-zA-Z0-9.\/]+)`)

// ConvertFlagParamsToRules converts received input rules parameters into a Firewall Profile rules array
func (fp *FirewallProfile) ConvertFlagParamsToRules(rulesIn string) error {
	for _, r := range strings.Split(rulesIn, ",") {
		values := firewallProfileRulesRegexp.FindStringSubmatch(r)
		if len(values) == 0 {
			return fmt.Errorf("invalid input rule format %s", r)
		}
		names := firewallProfileRulesRegexp.SubexpNames()
		mapValueByName := make(map[string]string)
		for j := range values {
			if names[j] != "" {
				mapValueByName[names[j]] = values[j]
			}
		}
		rule := new(Rule)
		// ip_protocol
		rule.Protocol = strings.ToLower(mapValueByName["ip_protocol"])

		// min_port
		minPort, err := strconv.Atoi(mapValueByName["min_port"])
		if err != nil {
			return fmt.Errorf("invalid port %s, is not a valid format. %s", mapValueByName["min_port"], err)
		}
		rule.MinPort = minPort

		// max_port
		maxPort, err := strconv.Atoi(mapValueByName["max_port"])
		if err != nil {
			maxPort = minPort
		}
		rule.MaxPort = maxPort

		// source/cidr
		rule.CidrIP = strings.ToLower(mapValueByName["source"])

		fp.Rules = append(fp.Rules, *rule)
	}
	return nil
}
