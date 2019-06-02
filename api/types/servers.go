package types

type Server struct {
	ID                string `json:"id" header:"ID"`
	Name              string `json:"name" header:"NAME"`
	Fqdn              string `json:"fqdn" header:"FQDN"`
	State             string `json:"state" header:"STATE"`
	PublicIP          string `json:"public_ip" header:"PUBLIC_IP"`
	PrivateIP         string `json:"private_ip,omitempty" header:"PRIVATE_IP"`
	TemplateID        string `json:"template_id" header:"TEMPLATE_ID"`
	ServerPlanID      string `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	CloudAccountID    string `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	SSHProfileID      string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
	FirewallProfileID string `json:"firewall_profile_id" header:"FIREWALL_PROFILE_ID"`
	ServerArrayID     string `json:"server_array_id,omitempty" header:"SERVER_ARRAY_ID" show:"nolist"`
	SubnetID          string `json:"subnet_id,omitempty" header:"SUBNET_ID" show:"nolist"`
	VpcID             string `json:"vpc_id,omitempty" header:"VPC_ID" show:"nolist"`
	Privateness       bool   `json:"privateness,omitempty" header:"PRIVATENESS" show:"nolist"`
	ResourceType      string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}

type ScriptChar struct {
	ResourceType    string                 `json:"resource_type" header:"RESOURCE_TYPE" show:"noshow,nolist"`
	ID              string                 `json:"id" header:"ID"`
	Type            string                 `json:"type" header:"TYPE"`
	ParameterValues map[string]interface{} `json:"parameter_values" header:"PARAMETER_VALUES"`
	TemplateID      string                 `json:"template_id" header:"TEMPLATE_ID"`
	ScriptID        string                 `json:"script_id" header:"SCRIPT_ID"`
	ExecutionOrder  int                    `json:"execution_order" header:"EXECUTION_ORDER" show:"noshow,nolist"`
}
