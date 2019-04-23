package testdata

import (
	"time"

	"github.com/ingrammicro/concerto/api/types"
)

// GetEventData loads test data
func GetEventData() []*types.Event {

	return []*types.Event{
		{
			ID:          "fakeID0",
			Timestamp:   time.Date(2014, 1, 1, 12, 0, 0, 0, time.UTC),
			Level:       "fakeLevel0",
			Header:      "fakeHeader0",
			Description: "fakeDescription0",
		},
		{
			ID:          "fakeID1",
			Timestamp:   time.Date(2015, 1, 10, 11, 0, 0, 0, time.UTC),
			Level:       "fakeLevel1",
			Header:      "fakeHeader1",
			Description: "fakeDescription1",
		},
	}
}
