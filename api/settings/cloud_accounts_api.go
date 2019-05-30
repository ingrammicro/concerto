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
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudAccountService{
		concertoService: concertoService,
	}, nil
}

// GetCloudAccountList returns the list of cloudAccounts as an array of CloudAccount
func (ca *CloudAccountService) GetCloudAccountList() (cloudAccounts []*types.CloudAccount, err error) {
	log.Debug("GetCloudAccountList")

	data, status, err := ca.concertoService.Get("/settings/cloud_accounts")
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

// GetCloudAccount returns a cloudAccount by its ID
func (ca *CloudAccountService) GetCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("GetCloudAccount")

	data, status, err := ca.concertoService.Get(fmt.Sprintf("/settings/cloud_accounts/%s", cloudAccountID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}
