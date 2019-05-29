package storage

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// StoragePlanService manages storage plan operations
type StoragePlanService struct {
	concertoService utils.ConcertoService
}

// NewStoragePlanService returns a Concerto storage plan service
func NewStoragePlanService(concertoService utils.ConcertoService) (*StoragePlanService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &StoragePlanService{
		concertoService: concertoService,
	}, nil
}

// GetStoragePlan returns a storage plan by its ID
func (sps *StoragePlanService) GetStoragePlan(storagePlanID string) (storagePlan *types.StoragePlan, err error) {
	log.Debug("GetStoragePlan")

	data, status, err := sps.concertoService.Get(fmt.Sprintf("/storage/plans/%s", storagePlanID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &storagePlan); err != nil {
		return nil, err
	}

	return storagePlan, nil
}
