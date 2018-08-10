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
