package types

import (
	"encoding/json"
)

// Template stores blueprint templates
type Template struct {
	ID                      string           `json:"id,omitempty" header:"ID"`
	Name                    string           `json:"name,omitempty" header:"NAME"`
	GenericImageID          string           `json:"generic_image_id,omitempty" header:"GENERIC_IMAGE_ID"`
	RunList                 []string         `json:"run_list,omitempty" header:"RUN_LIST" show:"nolist"`
	ConfigurationAttributes *json.RawMessage `json:"configuration_attributes,omitempty" header:"CONFIGURATION_ATTRIBUTES" show:"nolist"`
	ResourceType            string           `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	CookbookVersions        *json.RawMessage `json:"cookbook_versions,omitempty" header:"COOKBOOK_VERSIONS" show:"nolist"`
	State                   string           `json:"state" header:"STATE" show:"nolist"`
	LabelableFields
}

// TemplateScript stores a templates' script info
type TemplateScript struct {
	ID              string           `json:"id" header:"ID"`
	Type            string           `json:"type" header:"TYPE"`
	ExecutionOrder  int              `json:"execution_order" header:"EXECUTION_ORDER"`
	TemplateID      string           `json:"template_id" header:"TEMPLATE_ID"`
	ScriptID        string           `json:"script_id" header:"SCRIPT_ID"`
	ParameterValues *json.RawMessage `json:"parameter_values" header:"PARAMETER_VALUES"`
}

// TemplateServer stores servers associated with the template
type TemplateServer struct {
	ID                string `json:"id"  header:"ID"`
	Name              string `json:"name" header:"NAME"`
	Fqdn              string `json:"fqdn" header:"FQDN"`
	State             string `json:"state" header:"STATE"`
	PublicIP          string `json:"public_ip" header:"PUBLIC_IP"`
	TemplateID        string `json:"template_id" header:"TEMPLATE_ID"`
	ServerPlanID      string `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	SSHProfileID      string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
	FirewallProfileID string `json:"firewall_profile_id" header:"FIREWALL_PROFILE_ID"`
}

// TemplateScriptCredentials stores credentials to servers
type TemplateScriptCredentials interface{}
