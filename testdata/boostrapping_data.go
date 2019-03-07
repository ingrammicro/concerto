package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
	"encoding/json"
)

// GetBootstrappingConfigurationData loads test data
func GetBootstrappingConfigurationData() *types.BootstrappingConfiguration {

	attrs := json.RawMessage(`{"fakeAttribute0":"val0","fakeAttribute1":"val1"}`)
	test := types.BootstrappingConfiguration{
		Policyfiles:         []types.BootstrappingPolicyfile{
			{
				ID:          "fakeProfileID0",
				RevisionID:  "fakeProfileRevisionID0",
				DownloadURL: "fakeProfileDownloadURL0",
			},
			{
				ID:          "fakeProfileID1",
				RevisionID:  "fakeProfileRevisionID1",
				DownloadURL: "fakeProfileDownloadURL1",
			},
		},
		Attributes:          &attrs,
		AttributeRevisionID: "fakeAttributeRevisionID",
	}

	return &test
}

// GetBootstrappingContinuousReportData loads test data
func GetBootstrappingContinuousReportData() *types.BootstrappingContinuousReport{

	testBootstrappingContinuousReport := types.BootstrappingContinuousReport{
		Stdout: "Bootstrap log created",
	}

	return &testBootstrappingContinuousReport
}