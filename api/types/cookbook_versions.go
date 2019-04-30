package types

type CookbookVersion struct {
	ID                string   `json:"id" header:"ID"`
	Name              string   `json:"name" header:"NAME"`
	Version           string   `json:"version" header:"VERSION"`
	State             string   `json:"state" header:"STATE"`
	RevisionID        string   `json:"revision_id,omitempty" header:"REVISION_ID"`
	Description       string   `json:"description" header:"DESCRIPTION"`
	Recipes           []string `json:"recipes"  header:"RECIPES" show:"nolist"`
	ResourceType      string   `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
	PubliclyAvailable bool     `json:"publicly_available" header:"PUBLICLY_AVAILABLE" show:"nolist"`
	GlobalLegacy      bool     `json:"global_legacy" header:"GLOBAL_LEGACY" show:"nolist"`
	UploadURL         string   `json:"upload_url" header:"UPLOAD_URL" show:"noshow,nolist"`
	ErrorMessage      string   `json:"error_message" header:"ERROR_MESSAGE" show:"nolist"`
	LabelableFields
}
