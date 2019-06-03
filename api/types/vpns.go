package types

// Vpn stores an IMCO VPN item
type Vpn struct {
	ID           string   `json:"id" header:"ID"`
	State        string   `json:"state,omitempty" header:"STATE"`
	VpcID        string   `json:"vpc_id,omitempty" header:"VPC_ID"`
	VpnPlanID    string   `json:"vpn_plan_id,omitempty" header:"VPN_PLAN_ID"`
	PublicIP     string   `json:"public_ip,omitempty" header:"PUBLIC_IP"`
	ExposedCIDRs []string `json:"exposed_cidrs,omitempty" header:"EXPOSED_CIDRS"`
	ResourceType string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

// VpnPlan stores an IMCO VPN Plan item
type VpnPlan struct {
	ID           string `json:"id" header:"ID"`
	Name         string `json:"name,omitempty" header:"NAME"`
	Active       string `json:"active_active,omitempty" header:"ACTIVE"`
	RemoteID     string `json:"remote_id,omitempty" header:"REMOTE_ID"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
