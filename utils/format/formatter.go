package format

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

var osExit = os.Exit

// Formatter defines output printing interface
type Formatter interface {
	PrintItem(item interface{}) error
	PrintList(items interface{}) error
	//PrintList(items [][]string, headers []string) error
	PrintError(context string, err error)
	PrintFatal(context string, err error)
}

var formatter Formatter

// InitializeFormatter creates a singleton Formatter
func InitializeFormatter(ftype string, out io.Writer) {
	if ftype == "json" {
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
