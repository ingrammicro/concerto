package types

// Vpc stores an IMCO VPC item
type Vpc struct {
	ID                 string   `json:"id" header:"ID"`
	Name               string   `json:"name,omitempty" header:"NAME"`
	CIDR               string   `json:"cidr,omitempty" header:"CIDR"`
	State              string   `json:"state,omitempty" header:"STATE"`
	CloudAccountID     string   `json:"cloud_account_id,omitempty" header:"CLOUD_ACCOUNT_ID"`
	RealmProviderName  string   `json:"realm_provider_name,omitempty" header:"REALM_PROVIDER_NAME"`
	HasVPN             bool     `json:"has_vpn,omitempty" header:"HAS_VPN" show:"nolist,noshow"`
	AllowedSubnetTypes []string `json:"allowed_subnet_types,omitempty" header:"ALLOWED_SUBNET_TYPES"`
	ResourceType       string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	Brownfield         bool     `json:"brownfield,omitempty" header:"BROWNFIELD" show:"nolist,noshow"`
	LabelableFields
}
