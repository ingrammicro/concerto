package types

import (
	"time"
)

// Event stores an IMCO event item
type Event struct {
	ID          string    `json:"id" header:"ID"`
	Timestamp   time.Time `json:"timestamp" header:"TIMESTAMP"`
	Level       string    `json:"level" header:"LEVEL"`
	Header      string    `json:"header" header:"HEADER"`
	Description string    `json:"description" header:"DESCRIPTION"`
}
