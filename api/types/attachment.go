package types

// Attachment stores an IMCO Attachment item
type Attachment struct {
	ID           string `json:"id" header:"ID"`
	Name         string `json:"name,omitempty" header:"NAME"`
	Uploaded     bool   `json:"uploaded,omitempty" header:"UPLOADED"`
	UploadURL    string `json:"upload_url,omitempty" header:"UPLOAD_URL" show:"noshow,nolist"`
	DownloadURL  string `json:"download_url,omitempty" header:"DOWNLOAD_URL" show:"nolist"`
	ResourceType string `json:"resource_type" header:"RESOURCE_TYPE" show:"nolist"`
}
