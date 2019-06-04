package types

import "time"

type CloudAccount struct {
	ID                           string    `json:"id" header:"ID"`
	Name                         string    `json:"name" header:"NAME"`
	CloudProviderID              string    `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	CloudProviderName            string    `header:"CLOUD_PROVIDER_NAME"`
	SupportsImporting            bool      `json:"supports_importing" header:"SUPPORTS_IMPORTING" show:"nolist"`
	SupportsImportingVPCs        bool      `json:"supports_importing_vpcs" header:"SUPPORTS_IMPORTING_VPCS" show:"nolist"`
	SupportsImportingFloatingIPs bool      `json:"supports_importing_floating_ips" header:"SUPPORTS_IMPORTING_FLOATING_IPS" show:"nolist"`
	LastDiscoveredAt             time.Time `json:"last_discovered_at" header:"LAST_DISCOVERED" show:"nolist"`
	LastDiscoveredFailedAt       time.Time `json:"last_discovery_failed_at" header:"LAST_DISCOVERED_FAILED" show:"nolist"`
	ResourceType                 string    `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}

type RequiredCredentials interface{}
