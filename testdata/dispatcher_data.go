package testdata

import (
	"github.com/ingrammicro/concerto/api/types"
)

// GetScriptCharacterizationsData loads test data
func GetScriptCharacterizationsData() []*types.ScriptCharacterization {

	return []*types.ScriptCharacterization{
		{
			Script: types.DispatcherScript{
				Code:            "fakeCode1",
				UUID:            "fakeUUID1",
				AttachmentPaths: []string{"fakeAttachmentPath1"},
			},
			UUID:       "fakeUUID1",
			Order:      0,
			Parameters: map[string]string{"fakeParamKey1": "fakeParamValue1"},
		},
	}
}

// GetScriptConclusionData loads test data
func GetScriptConclusionData() *types.ScriptConclusion {

	return &types.ScriptConclusion{
		Output:     "fakeOutput1",
		ExitCode:   0,
		StartedAt:  "fakeStartedAt1",
		FinishedAt: "fakeFinishedAt1",
	}
}

// GetDownloadAttachmentData loads test data
func GetDownloadAttachmentData() map[string]string {
	return map[string]string{
		"fakeEndpoint":      "/blueprint/attachments/fakeID1",
		"fakeAttachmentDir": "/tmp/fakeFolderID1/attachments",
	}
}
