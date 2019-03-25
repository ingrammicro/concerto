package format

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

// Required workaround for testing os.Exit(1) scenarios in Go with coverage.
// Otherwise, PrintFatal cannot be evaluated due to os.Exit() cannot be captured.
// Implemented in test files (json/text)
var osExit = os.Exit

// Formatter defines output printing interface
type Formatter interface {
	PrintItem(item interface{}) error
	PrintList(items interface{}) error
	PrintError(context string, err error)
	PrintFatal(context string, err error)
}

var formatter Formatter

// InitializeFormatter creates a singleton Formatter
func InitializeFormatter(formatterType string, out io.Writer) {
	if formatterType == "json" {
		formatter = NewJSONFormatter(out)
	} else {
		formatter = NewTextFormatter(out)
	}
}

// GetFormatter creates a new JSONFormatter
func GetFormatter() Formatter {
	if formatter != nil {
		return formatter
	}
	log.Warn("Formatter hasn't been initialized. Initializing now to default formatter")
	InitializeFormatter("", os.Stdout)
	return formatter
}
