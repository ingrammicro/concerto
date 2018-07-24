package types

// CloudProvider stores Cloud Provider data
type CloudProvider struct {
	ID   string `json:"id" header:"ID"`
	Name string `json:"name" header:"NAME"`
}
