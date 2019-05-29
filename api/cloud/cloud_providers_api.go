package cloud

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// CloudProviderService manages cloudProvider operations
type CloudProviderService struct {
	concertoService utils.ConcertoService
}

// NewCloudProviderService returns a Concerto cloudProvider service
func NewCloudProviderService(concertoService utils.ConcertoService) (*CloudProviderService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudProviderService{
		concertoService: concertoService,
	}, nil
}

// GetCloudProviderList returns the list of cloudProviders as an array of CloudProvider
func (cl *CloudProviderService) GetCloudProviderList() (cloudProviders []*types.CloudProvider, err error) {
	log.Debug("GetCloudProviderList")

	data, status, err := cl.concertoService.Get("/cloud/cloud_providers")
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

// GetServerStoragePlanList returns the list of storage plans as an array of StoragePlan
func (dm *CloudProviderService) GetServerStoragePlanList(providerID string) (storagePlans []*types.StoragePlan, err error) {
	log.Debug("GetServerStoragePlanList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &storagePlans); err != nil {
		return nil, err
	}

	return storagePlans, nil
}
