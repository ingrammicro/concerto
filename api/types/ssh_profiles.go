package types

type SSHProfile struct {
	ID         string `json:"id" header:"ID"`
	Name       string `json:"name" heade:"NAME"`
	PublicKey  string `json:"public_key" header:"PUBLIC_KEY"`
	PrivateKey string `json:"private_key" header:"PRIVATE_KEY"`
}
