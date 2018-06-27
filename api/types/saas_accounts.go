package types

type SaasAccount struct {
	ID             string `json:"id" header:"ID"`
	SaasProviderID string `json:"saas_provider_id" header:"SAAS PROVIDER ID"`
}

type SaasRequiredCredentials interface{}
