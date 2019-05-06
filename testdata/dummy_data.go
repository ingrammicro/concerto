package testdata

import "time"

// DummyStructTestFormatter resolves type for testing
type DummyStructTestFormatter struct {
	ID               string                 `json:"id" header:"ID"`
	RemainingSeconds float32                `json:"remaining_seconds" header:"REMAINING SECONDS" show:"minifySeconds"`
	JSONRaw          map[string]interface{} `json:"json_raw" header:"JSON RAW"`
	Time             time.Time              `json:"time" header:"TIME"`
}

// GetDummyData loads test data
func GetDummyData() []*DummyStructTestFormatter {

	return []*DummyStructTestFormatter{
		{
			ID:               "fakeID0",
			RemainingSeconds: 100360012,
			JSONRaw:          map[string]interface{}{"fakeFlavour01": "x", "fakeFlavour02": "y"},
			Time:             time.Now(),
		},
		{
			ID:               "fakeID1",
			RemainingSeconds: 1,
			JSONRaw:          map[string]interface{}{"fakeFlavour11": "a", "fakeFlavour12": "b"},
			Time:             time.Now(),
		},
	}
}
