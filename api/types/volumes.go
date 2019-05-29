package types

type Volume struct {
	ID               string `json:"id" header:"ID"`
	Name             string `json:"name" header:"NAME"`
	Size             int    `json:"size" header:"SIZE"`
	State            string `json:"state" header:"STATE"`
	Device           string `json:"device" header:"DEVICE"`
	StoragePlanID    string `json:"storage_plan_id,omitempty" header:"STORAGE_PLAN_ID"`
	CloudAccountID   string `json:"cloud_account_id,omitempty" header:"CLOUD_ACCOUNT_ID"`
	RealmID          string `json:"realm_id,omitempty" header:"REALM_ID"`
	AttachedServerID string `json:"attached_server_id,omitempty" header:"ATTACHED_SERVER_ID"`
	Brownfield       bool   `json:"brownfield,omitempty" header:"BROWNFIELD" show:"nolist,noshow"`
	ResourceType     string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelableFields
}
