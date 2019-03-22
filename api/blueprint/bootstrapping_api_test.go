package blueprint

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBootstrappingServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewBootstrappingService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetBootstrappingConfiguration(t *testing.T) {
	bcIn := testdata.GetBootstrappingConfigurationData()
	GetBootstrappingConfigurationMocked(t, bcIn)
	GetBootstrappingConfigurationFailErrMocked(t, bcIn)
	GetBootstrappingConfigurationFailStatusMocked(t, bcIn)
	GetBootstrappingConfigurationFailJSONMocked(t, bcIn)
}

func TestReportBootstrappingAppliedConfiguration(t *testing.T) {
	bcIn := testdata.GetBootstrappingConfigurationData()
	ReportBootstrappingAppliedConfigurationMocked(t, bcIn)
	ReportBootstrappingAppliedConfigurationFailErrMocked(t, bcIn)
	ReportBootstrappingAppliedConfigurationFailStatusMocked(t, bcIn)
	ReportBootstrappingAppliedConfigurationFailJSONMocked(t, bcIn)
}

func TestReportBootstrappingLog(t *testing.T) {
	commandIn := testdata.GetBootstrappingContinuousReportData()
	ReportBootstrappingLogMocked(t, commandIn)
	ReportBootstrappingLogFailErrMocked(t, commandIn)
	ReportBootstrappingLogFailStatusMocked(t, commandIn)
	ReportBootstrappingLogFailJSONMocked(t, commandIn)
}

func TestDownloadPolicyfile(t *testing.T) {
	dataIn := testdata.GetBootstrappingDownloadFileData()
	DownloadPolicyfileMocked(t, dataIn)
	DownloadPolicyfileFailErrMocked(t, dataIn)
}
