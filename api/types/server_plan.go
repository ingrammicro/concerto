package types

type ServerPlan struct {
	ID                string  `json:"id" header:"ID"`
	Name              string  `json:"name" header:"NAME"`
	Memory            int     `json:"memory" header:"MEMORY"`
	CPUs              float32 `json:"cpus" header:"CPUS"`
	Storage           int     `json:"storage" header:"STORAGE"`
	LocationID        string  `json:"location_id" header:"LOCATION_ID"`
	LocationName      string  `json:"location_name" header:"LOCATION_NAME"`
	CloudProviderID   string  `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName string  `json:"cloud_provider_name" header:"CLOUD_PROVIDER_NAME"`
}
