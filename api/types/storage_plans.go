package types

type StoragePlan struct {
	ID                  string `json:"id" header:"ID"`
	Name                string `json:"name" header:"NAME"`
	MinSize             int    `json:"min_size,omitempty" header:"MIN_SIZE"`
	MaxSize             int    `json:"max_size,omitempty" header:"MAX_SIZE"`
	CloudProviderID     string `json:"cloud_provider_id,omitempty" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName   string `json:"cloud_provider_name" header:"CLOUD_PROVIDER_NAME"`
	LocationID          string `json:"location_id,omitempty" header:"LOCATION_ID"`
	LocationName        string `json:"location_name" header:"LOCATION_NAME"`
	RealmID             string `json:"realm_id,omitempty" header:"REALM_ID"`
	RealmProviderName   string `json:"realm_provider_name,omitempty" header:"REALM_PROVIDER_NAME"`
	FlavourProviderName string `json:"flavour_provider_name,omitempty" header:"FLAVOUR_PROVIDER_NAME"`
	Deprecated          bool   `json:"deprecated,omitempty" header:"DEPRECATED" show:"nolist,noshow"`
	ResourceType        string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
