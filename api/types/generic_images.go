package types

// GenericImage stores Generic Image data
type GenericImage struct {
	ID   string `json:"id" header:"ID"`
	Name string `json:"name" header:"NAME"`
}
