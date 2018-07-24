package types

// SaasProvider stores Saas Provider data
type SaasProvider struct {
	ID                  string   `json:"id" header:"ID"`
	Name                string   `json:"name" header:"NAME"`
	RequiredAccountData []string `json:"required_account_data" header:"REQUIRED_ACCOUNT_DATA"`
}
