package testdata

import (
	"encoding/json"
)

// DummyStructTestFormatter resolves type for testing
type DummyStructTestFormatter struct {
	ID               string           `json:"id" header:"ID"`
	RemainingSeconds float32          `json:"remaining_seconds" header:"REMAINING SECONDS" show:"minifySeconds"`
	JSONRaw          json.RawMessage  `json:"json_raw" header:"JSON RAW"`
	JSONRawPtr       *json.RawMessage `json:"json_raw_ptr" header:"JSON RAW PTR"`
}

// GetDummyData loads test data
func GetDummyData() []*DummyStructTestFormatter {

	param0 := json.RawMessage(`{"fakeFlavour01":"x","fakeFlavour02":"y"}`)
	param1 := json.RawMessage(`{"fakeFlavour11":"a","fakeFlavour12":"b"}`)

	return []*DummyStructTestFormatter{
		{
			ID:               "fakeID0",
			RemainingSeconds: 100360012,
			JSONRaw:          param0,
			JSONRawPtr:       &param1,
		},
		{
			ID:               "fakeID1",
			RemainingSeconds: 1,
			JSONRaw:          param0,
			JSONRawPtr:       nil,
		},
	}
}
