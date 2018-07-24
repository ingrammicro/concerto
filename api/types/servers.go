package types

// Server stores Server data
type Server struct {
	ID                string `json:"id" header:"ID"`
	Name              string `json:"name" header:"NAME"`
	Fqdn              string `json:"fqdn" header:"FQDN"`
	State             string `json:"state" header:"STATE"`
	PublicIP          string `json:"public_ip" header:"PUBLIC_IP"`
	TemplateID        string `json:"template_id" header:"TEMPLATE_ID"`
	ServerPlanID      string `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	CloudAccountID    string `json:"cloud_account_id" header:"CLOUD_ACCOUNT_ID"`
	SSHProfileID      string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
	FirewallProfileID string `json:"firewall_profile_id" header:"FIREWALL_PROFILE_ID"`
}

// DNS stores DNS data
type DNS struct {
	ID       string `json:"id" header:"ID"`
	Name     string `json:"name" header:"NAME"`
	Content  string `json:"content" header:"CONTENT"`
	Type     string `json:"type" header:"TYPE"`
	IsFQDN   bool   `json:"is_fqdn" header:"IS_FQDN"`
	DomainID string `json:"domain_id" header:"DOMAIN_ID"`
}

// ScriptChar stores Script Char data
type ScriptChar struct {
	ID              string   `json:"id" header:"ID"`
	Type            string   `json:"type" header:"TYPE"`
	ParameterValues struct{} `json:"parameter_values" header:"PARAMETER_VALUES"`
	TemplateID      string   `json:"template_id" header:"TEMPLATE_ID"`
	ScriptID        string   `json:"script_id" header:"SCRIPT_ID"`
}
