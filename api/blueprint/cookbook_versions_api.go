package blueprint

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// CookbookVersionService manages cookbook version operations
type CookbookVersionService struct {
	concertoService utils.ConcertoService
}

// NewCookbookVersionService returns a Concerto cookbook version service
func NewCookbookVersionService(concertoService utils.ConcertoService) (*CookbookVersionService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CookbookVersionService{
		concertoService: concertoService,
	}, nil
}

// GetCookbookVersionList returns the list of cookbook versions as an array of CookbookVersion
func (cv *CookbookVersionService) GetCookbookVersionList() (cookbookVersions []*types.CookbookVersion, err error) {
	log.Debug("GetCookbookVersionList")

	data, status, err := cv.concertoService.Get("/blueprint/cookbook_versions")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersions); err != nil {
		return nil, err
	}

	return cookbookVersions, nil
}

// GetCookbookVersion returns a cookbook version by its ID
func (cv *CookbookVersionService) GetCookbookVersion(ID string) (cookbookVersion *types.CookbookVersion, err error) {
	log.Debug("GetCookbookVersion")

	data, status, err := cv.concertoService.Get(fmt.Sprintf("/blueprint/cookbook_versions/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersion); err != nil {
		return nil, err
	}

	return cookbookVersion, nil
}

// CreateCookbookVersion creates a new cookbook version
func (cv *CookbookVersionService) CreateCookbookVersion(cvVector *map[string]interface{}) (cookbookVersion *types.CookbookVersion, err error) {
	log.Debug("CreateCookbookVersion")

	data, status, err := cv.concertoService.Post("/blueprint/cookbook_versions", cvVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersion); err != nil {
		return nil, err
	}

	return cookbookVersion, nil
}

// UploadCookbookVersion uploads a cookbook version file
func (cv *CookbookVersionService) UploadCookbookVersion(sourceFilePath string, targetURL string) error {
	log.Debug("UploadCookbookVersion")

	data, status, err := cv.concertoService.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ProcessCookbookVersion process a cookbook version by its ID
func (cv *CookbookVersionService) ProcessCookbookVersion(cvVector *map[string]interface{}, ID string) (cookbookVersion *types.CookbookVersion, err error) {
	log.Debug("ProcessCookbookVersion")

	data, status, err := cv.concertoService.Post(fmt.Sprintf("/blueprint/cookbook_versions/%s/process", ID), cvVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersion); err != nil {
		return nil, err
	}

	return cookbookVersion, nil
}

// DeleteCookbookVersion deletes a cookbook version by its ID
func (cv *CookbookVersionService) DeleteCookbookVersion(ID string) (err error) {
	log.Debug("DeleteCookbookVersion")

	data, status, err := cv.concertoService.Delete(fmt.Sprintf("/blueprint/cookbook_versions/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
