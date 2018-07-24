package types

// Location stores Location data
type Location struct {
	ID   string `json:"id" header:"ID"`
	Name string `json:"name" header:"NAME"`
}
