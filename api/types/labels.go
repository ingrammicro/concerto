package types

type Label struct {
	ID           string `json:"id" header:"ID"`
	Name         string `json:"name" header:"NAME"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE"`
	Namespace    string `json:"namespace" header:"NAMESPACE" show:"nolist"`
	Value        string `json:"value" header:"VALUE" show:"nolist"`
}

type LabeledResource struct {
	ID           string `json:"id" header:"ID"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE"`
}

type LabelableFields struct {
	LabelIDs []string `json:"label_ids" header:"LABEL_IDS" show:"nolist,noshow"`
	Labels   []string `json:"labels" header:"LABELS"`
}

type Labelable interface {
	FilterByLabelIDs(labelIDs []string) bool
	AssignLabelIDs(labelIDs []string)
	FillInLabelNames(labelNamesByID map[string]string)
}

func (lf *LabelableFields) FilterByLabelIDs(labelIDs []string) bool {
	for _, lid := range labelIDs {
		var labelIDFound bool
		for _, resourceLabelID := range lf.LabelIDs {
			if lid == resourceLabelID {
				labelIDFound = true
				break
			}
		}
		if !labelIDFound {
			return false
		}
	}
	return true
}

func (lf *LabelableFields) AssignLabelIDs(labelIDs []string) {
	for _, lid := range labelIDs {
		for _, resourceLabelID := range lf.LabelIDs {
			if lid == resourceLabelID {
				break
			}
		}
	}
}

func (lf *LabelableFields) FillInLabelNames(labelNamesByID map[string]string) {
	for lID, lName := range labelNamesByID {
		for _, resourceLabelID := range lf.LabelIDs {
			if lID == resourceLabelID {
				lf.Labels = append(lf.Labels, lName)
			}
		}
	}
}
