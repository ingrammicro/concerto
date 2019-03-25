package types

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
