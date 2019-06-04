package types

type ServerArray struct {
	ID                string `json:"id" header:"ID"`
	Name              string `json:"name" header:"NAME"`
	State             string `json:"state" header:"STATE"`
	Size              int    `json:"size" header:"SIZE"`
	TemplateID        string `json:"template_id" header:"TEMPLATE_ID"`
	CloudAccountID    string `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	ServerPlanID      string `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	FirewallProfileID string `json:"firewall_profile_id" header:"FIREWALL_PROFILE_ID"`
	SSHProfileID      string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
	SubnetID          string `json:"subnet_id,omitempty" header:"SUBNET_ID" show:"nolist"`
	VpcID             string `json:"vpc_id,omitempty" header:"VPC_ID" show:"nolist"`
	Privateness       bool   `json:"privateness,omitempty" header:"PRIVATENESS" show:"nolist"`
	ResourceType      string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}
