package types

type CloudAccount struct {
	Id            string `json:"id" header:"ID"`
	Name          string `json:"name" header:"NAME"`
	CloudProvId   string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	CloudProvName string `json:"cloud_provider_name" header:"CLOUD_PROVIDER_NAME"`
}

type RequiredCredentials interface{}
