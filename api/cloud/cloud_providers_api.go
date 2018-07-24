package cloud

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// CloudProviderService manages Cloud Provider Service operations
type CloudProviderService struct {
	concertoService utils.ConcertoService
}

// NewCloudProviderService returns a Concerto cloudProvider service
func NewCloudProviderService(concertoService utils.ConcertoService) (*CloudProviderService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &CloudProviderService{
		concertoService: concertoService,
	}, nil
}

// GetCloudProviderList returns the list of cloudProviders as an array of CloudProvider
func (cl *CloudProviderService) GetCloudProviderList() (cloudProviders []types.CloudProvider, err error) {
	log.Debug("GetCloudProviderList")

	data, status, err := cl.concertoService.Get("/v1/cloud/cloud_providers")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudProviders); err != nil {
		return nil, err
	}

	return cloudProviders, nil
}
