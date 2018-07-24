package types

// SaasAccount stores Saas Account data
type SaasAccount struct {
	ID             string `json:"id" header:"ID"`
	SaasProviderID string `json:"saas_provider_id" header:"SAAS PROVIDER ID"`
}

// SaasRequiredCredentials stores Saas Required Credentials data
type SaasRequiredCredentials interface{}
