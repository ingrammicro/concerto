package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetPollingPingData loads test data
func GetPollingPingData() *types.PollingPing {

	testPollingPing := types.PollingPing{
		PendingCommands: true,
	}

	return &testPollingPing
}

// GetPollingCommandData loads test data
func GetPollingCommandData() *types.PollingCommand {

	testPollingCommand := types.PollingCommand{
		Id:       "fakeID0",
		Script:   "fakeScript",
		Stdout:   "fakeStdout",
		Stderr:   "fakeStdin",
		ExitCode: 0,
	}

	return &testPollingCommand
}

// GetPollingContinuousReportData loads test data
func GetPollingContinuousReportData() *types.PollingContinuousReport {

	testPollingContinuousReport := types.PollingContinuousReport{
		Stdout: "Bootstrap log created",
	}

	return &testPollingContinuousReport
}
