package types

// FloatingIP stores an IMCO Floating IP item
type FloatingIP struct {
	ID               string `json:"id" header:"ID"`
	Name             string `json:"name,omitempty" header:"NAME"`
	Address          string `json:"address,omitempty" header:"ADDRESS"`
	State            string `json:"state,omitempty" header:"STATE"`
	CloudAccountID   string `json:"cloud_account_id,omitempty" header:"CLOUD_ACCOUNT_ID"`
	RealmID          string `json:"realm_id,omitempty" header:"REALM_ID"`
	AttachedServerID string `json:"attached_server_id,omitempty" header:"ATTACHED_SERVER_ID"`
	ResourceType     string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	Brownfield       bool   `json:"brownfield,omitempty" header:"BROWNFIELD" show:"nolist,noshow"`
	LabelableFields
}
