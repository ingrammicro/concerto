package types

type SSHProfile struct {
	ID           string   `json:"id" header:"ID"`
	Name         string   `json:"name" header:"NAME"`
	PublicKey    string   `json:"public_key" header:"PUBLIC_KEY"`
	PrivateKey   string   `json:"private_key" header:"PRIVATE_KEY" show:"nolist"`
	ResourceType string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelIDs     []string `json:"label_ids" header:"LABEL_IDS" show:"nolist,noshow"`
	Labels       []string `json:"labels" header:"LABELS"`
}
