package types

type WizardApp struct {
	ID                  string                 `json:"id" header:"ID"`
	Name                string                 `json:"name" header:"NAME"`
	FlavourRequirements map[string]interface{} `json:"flavour_requirements" header:"FLAVOUR_REQUIREMENTS"`
	GenericImageID      string                 `json:"generic_image_id" header:"GENERIC_IMAGE_ID"`
}
