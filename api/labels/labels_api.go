package labels

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// LabelService manages polling operations
type LabelService struct {
	concertoService utils.ConcertoService
}

// NewLabelService returns a Concerto labels service
func NewLabelService(concertoService utils.ConcertoService) (*LabelService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &LabelService{
		concertoService: concertoService,
	}, nil
}

// GetLabelList returns the list of labels as an array of Label
func (lbl *LabelService) GetLabelList() (labels []types.Label, err error) {
	log.Debug("GetLabelList")

	data, status, err := lbl.concertoService.Get("/v1/labels")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &labels); err != nil {
		return nil, err
	}

	// exclude internal labels (with a Namespace defined)
	var filteredLabels []types.Label
	for _, label := range labels {
		if label.Namespace == "" {
			filteredLabels = append(filteredLabels, label)
		}
	}

	return filteredLabels, nil
}

// CreateLabel creates a label
func (lbl *LabelService) CreateLabel(labelVector *map[string]interface{}) (label *types.Label, err error) {
	log.Debug("CreateLabel")

	data, status, err := lbl.concertoService.Post("/v1/labels/", labelVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &label); err != nil {
		return nil, err
	}

	return label, nil
}

// AddLabel assigns a single label from a single labelable resource
func (lbl *LabelService) AddLabel(labelVector *map[string]interface{}, labelID string) (labeledResources []types.LabeledResources, err error) {
	log.Debug("AddLabel")

	data, status, err := lbl.concertoService.Post(fmt.Sprintf("/v1/labels/%s/resources", labelID), labelVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &labeledResources); err != nil {
		return nil, err
	}

	return labeledResources, nil
}

// RemoveLabel de-assigns a single label from a single labelable resource
func (lbl *LabelService) RemoveLabel(labelID string, resourceType string, resourceID string) error {
	log.Debug("RemoveLabel")

	data, status, err := lbl.concertoService.Delete(fmt.Sprintf("v1/labels/%s/resources/%s/%s", labelID, resourceType, resourceID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
