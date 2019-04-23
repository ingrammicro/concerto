package cloud

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// GenericImageService manages genericImage operations
type GenericImageService struct {
	concertoService utils.ConcertoService
}

// NewGenericImageService returns a Concerto genericImage service
func NewGenericImageService(concertoService utils.ConcertoService) (*GenericImageService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &GenericImageService{
		concertoService: concertoService,
	}, nil
}

// GetGenericImageList returns the list of generic images as an array of GenericImage
func (cl *GenericImageService) GetGenericImageList() (genericImages []*types.GenericImage, err error) {
	log.Debug("GetGenericImageList")

	data, status, err := cl.concertoService.Get("/cloud/generic_images")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &genericImages); err != nil {
		return nil, err
	}

	return genericImages, nil
}
