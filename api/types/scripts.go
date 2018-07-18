package types

// Script holds script data
type Script struct {
	ID           string   `json:"id" header:"ID"`
	Name         string   `json:"name" header:"NAME"`
	Description  string   `json:"description" header:"DESCRIPTION"`
	Code         string   `json:"code" header:"CODE" show:"nolist"`
	Parameters   []string `json:"parameters" header:"PARAMETERS"`
	ResourceType string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	LabelIDs     []string `json:"label_ids" header:"LABEL_IDS" show:"nolist,noshow"`
	Labels       []string `json:"labels" header:"LABELS"`
}
