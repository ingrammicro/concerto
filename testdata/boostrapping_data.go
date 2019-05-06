package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetBootstrappingConfigurationData loads test data
func GetBootstrappingConfigurationData() *types.BootstrappingConfiguration {

	return &types.BootstrappingConfiguration{
		Policyfiles: []types.BootstrappingPolicyfile{
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
		Attributes:          map[string]interface{}{"fakeAttribute0": "val0", "fakeAttribute1": "val1"},
		AttributeRevisionID: "fakeAttributeRevisionID",
	}
}

// GetBootstrappingContinuousReportData loads test data
func GetBootstrappingContinuousReportData() *types.BootstrappingContinuousReport {

	return &types.BootstrappingContinuousReport{
		Stdout: "Bootstrap log created",
	}
}

//GetBootstrappingDownloadFileData loads test data
func GetBootstrappingDownloadFileData() map[string]string {
	return map[string]string{
		"fakeURLToFile":        "http://fakeURLToFile.xxx/filename.tgz",
		"fakeFileDownloadFile": "filename.tgz",
	}
}
