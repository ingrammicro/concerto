package types

// LoadBalancer stores Load Balancer data
type LoadBalancer struct {
	ID                       string `json:"id" header:"ID"`
	Name                     string `json:"name" header:"NAME"`
	Fqdn                     string `json:"fqdn" header:"FQDN"`
	Protocol                 string `json:"protocol" header:"PROTOCOL"`
	Port                     int    `json:"port" header:"PORT"`
	Algorithm                string `json:"algorithm" header:"ALGORITHM"`
	SSLCertificate           string `json:"ssl_certificate" header:"SSL_CERTIFICATE"`
	SSLCertificatePrivateKey string `json:"ssl_certificate_private_key" header:"SSL_CERTIFICATE_PRIVATE_KEY"`
	DomainID                 string `json:"domain_id" header:"DOMAIN_ID"`
	CloudProviderID          string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	TrafficIn                int    `json:"traffic_in" header:"TRAFFIC_IN"`
	TrafficOut               int    `json:"traffic_out" header:"TRAFFIC_OUT"`
}

// LBNode stores Load Balancer Node data
type LBNode struct {
	ID       string `json:"id" header:"ID"`
	Name     string `json:"name" header:"NAME"`
	PublicIP string `json:"public_ip" header:"PUBLIC_IP"`
	State    string `json:"state" header:"STATE"`
	ServerID string `json:"server_id" header:"SERVER_ID"`
	Port     int    `json:"port" header:"PORT"`
}
