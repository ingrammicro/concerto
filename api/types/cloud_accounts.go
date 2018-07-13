package types

type CloudAccount struct {
	ID                string `json:"id" header:"ID"`
	Name              string `json:"name" header:"NAME"`
	CloudProviderID   string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName string `json:"cloud_provider_name" header:"CLOUD_PROVIDER_NAME"`
}

type RequiredCredentials interface{}
