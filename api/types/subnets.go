package types

// Subnet stores an IMCO VPC Subnet item
type Subnet struct {
	ID                     string `json:"id" header:"ID"`
	Name                   string `json:"name,omitempty" header:"NAME"`
	CIDR                   string `json:"cidr,omitempty" header:"CIDR"`
	State                  string `json:"state,omitempty" header:"STATE"`
	Type                   string `json:"type,omitempty" header:"STATE"`
	VpcID                  string `json:"vpc_id,omitempty" header:"VPC_ID"`
	ServerCreationDisabled bool   `json:"server_creation_disabled,omitempty" header:"SERVER_CREATION_DISABLED" show:"nolist,noshow"`
	ResourceType           string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	Brownfield             bool   `json:"brownfield,omitempty" header:"BROWNFIELD" show:"nolist,noshow"`
}
