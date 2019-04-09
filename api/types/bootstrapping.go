package types

import (
	"encoding/json"
)

type BootstrappingConfiguration struct {
	Policyfiles         []BootstrappingPolicyfile `json:"policyfiles,omitempty" header:"POLICY_FILES" show:"nolist"`
	Attributes          *json.RawMessage          `json:"attributes,omitempty" header:"ATTRIBUTES" show:"nolist"`
	AttributeRevisionID string                    `json:"attribute_revision_id,omitempty" header:"ATTRIBUTE_REVISION_ID"`
}

type BootstrappingPolicyfile struct {
	ID          string `json:"id,omitempty" header:"ID"`
	RevisionID  string `json:"revision_id,omitempty" header:"REVISION_ID"`
	DownloadURL string `json:"download_url,omitempty" header:"DOWNLOAD_URL"`
}

type BootstrappingContinuousReport struct {
	Stdout string `json:"stdout" header:"STDOUT"`
}

type BootstrappingAppliedConfiguration struct {
	StartedAt             string `json:"started_at,omitempty" header:"STARTED_AT"`
	FinishedAt            string `json:"finished_at,omitempty" header:"FINISHED_AT"`
	PolicyfileRevisionIDs string `json:"policyfile_revision_ids,omitempty" header:"POLICY_FILE_REVISION_IDS" show:"nolist"`
	AttributeRevisionID   string `json:"attribute_revision_id,omitempty" header:"ATTRIBUTE_REVISION_ID"`
}
