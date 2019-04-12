package types

type Firewall struct {
	Profile Policy `json:"firewall_profile"`
}

type Policy struct {
	Rules       []PolicyRule `json:"rules"`
	Md5         string       `json:"md5,omitempty"`
	ActualRules []PolicyRule `json:"actual_rules,omitempty"`
}

type PolicyRule struct {
	Name     string `json:"name,omitempty" header:"NAME" show:"nolist"`
	Cidr     string `json:"cidr_ip" header:"CIDR"`
	Protocol string `json:"ip_protocol" header:"PROTOCOL"`
	MinPort  int    `json:"min_port" header:"MIN"`
	MaxPort  int    `json:"max_port" header:"MAX"`
}

// CheckPolicyRule checks if rule belongs to Policy
func (p *Policy) CheckPolicyRule(rule PolicyRule) bool {
	exists := false
	for _, policyRule := range p.Rules {
		if (policyRule.Cidr == rule.Cidr) && (policyRule.MaxPort == rule.MaxPort) && (policyRule.MinPort == rule.MinPort) && (policyRule.Protocol == rule.Protocol) {
			exists = true
		}
	}
	return exists
}
