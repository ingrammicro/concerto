package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetPollingPingData loads test data
func GetPollingPingData() *types.PollingPing {

	return &types.PollingPing{
		PendingCommands: true,
	}
}

// GetPollingCommandData loads test data
func GetPollingCommandData() *types.PollingCommand {

	return &types.PollingCommand{
		ID:       "fakeID0",
		Script:   "fakeScript0",
		Stdout:   "fakeStdout0",
		Stderr:   "fakeStderr0",
		ExitCode: 0,
	}
}

// GetPollingContinuousReportData loads test data
func GetPollingContinuousReportData() *types.PollingContinuousReport {

	return &types.PollingContinuousReport{
		Stdout: "Bootstrap log created",
	}
}
