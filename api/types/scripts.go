package types

// Script stores Script data
type Script struct {
	ID          string   `json:"id" header:"ID"`
	Name        string   `json:"name" header:"NAME"`
	Description string   `json:"description" header:"DESCRIPTION"`
	Code        string   `json:"code" header:"CODE" show:"nolist"`
	Parameters  []string `json:"parameters" header:"PARAMETERS"`
}
