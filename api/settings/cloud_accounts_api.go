package settings

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// CloudAccountService manages cloudAccount operations
type CloudAccountService struct {
	concertoService utils.ConcertoService
}

// NewCloudAccountService returns a Concerto cloudAccount service
func NewCloudAccountService(concertoService utils.ConcertoService) (*CloudAccountService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &CloudAccountService{
		concertoService: concertoService,
	}, nil
}

// GetCloudAccountList returns the list of cloudAccounts as an array of CloudAccount
func (ca *CloudAccountService) GetCloudAccountList() (cloudAccounts []types.CloudAccount, err error) {
	log.Debug("GetCloudAccountList")

	data, status, err := ca.concertoService.Get("/v2/settings/cloud_accounts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccounts); err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}
