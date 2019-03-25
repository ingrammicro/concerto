package types

import (
	"encoding/json"
)

type BootstrappingConfiguration struct {
	Policyfiles         []BootstrappingPolicyfile `json:"policyfiles,omitempty" header:"POLICY FILES" show:"nolist"`
	Attributes          *json.RawMessage          `json:"attributes,omitempty" header:"ATTRIBUTES" show:"nolist"`
	AttributeRevisionID string                    `json:"attribute_revision_id,omitempty" header:"ATTRIBUTE REVISION ID"`
}

type BootstrappingPolicyfile struct {
	ID          string `json:"id,omitempty" header:"ID"`
	RevisionID  string `json:"revision_id,omitempty" header:"REVISION ID"`
	DownloadURL string `json:"download_url,omitempty" header:"DOWNLOAD URL"`
}

type BootstrappingContinuousReport struct {
	Stdout string `json:"stdout" header:"STDOUT"`
}

type BootstrappingAppliedConfiguration struct {
	StartedAt             string `json:"started_at,omitempty" header:"STARTED AT"`
	FinishedAt            string `json:"finished_at,omitempty" header:"FINISHED AT"`
	PolicyfileRevisionIDs string `json:"policyfile_revision_ids,omitempty" header:"POLICY FILE REVISION IDS" show:"nolist"`
	AttributeRevisionID   string `json:"attribute_revision_id,omitempty" header:"ATTRIBUTE REVISION ID"`
}
