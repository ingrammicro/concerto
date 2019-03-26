package format

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ingrammicro/concerto/api/cloud"
	"testing"

	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPrintItemJSON(t *testing.T) {

	assert := assert.New(t)
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {

		serverOut := cloud.GetServerMocked(t, &serverIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("json", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*serverOut)
		assert.Nil(err, "JSON formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^\\{\\\"id\\\":.*\\}", b.String(), "JSON output didn't match regular expression")
	}
}

func TestPrintItemTemplateJSON(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {

		templateOut := blueprint.GetTemplateMocked(t, &templateIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("json", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*templateOut)
		assert.Nil(err, "JSON formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^\\{\\\"id\\\":.*\\}", b.String(), "JSON output didn't match regular expression")
	}
}

func TestPrintListJSON(t *testing.T) {

	assert := assert.New(t)
	serversIn := testdata.GetServerData()
	serverOut := cloud.GetServerListMocked(t, serversIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(serverOut)
	assert.Nil(err, "JSON formatter PrintList error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp("^\\[\\{\\\"id\\\":.*\\}\\]", b.String(), "JSON output didn't match regular expression")
}

func TestPrintListTemplateJSON(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	templatesOut := blueprint.GetTemplateListMocked(t, templatesIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(*templatesOut)
	assert.Nil(err, "JSON formatter PrintList error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp("^\\[\\{\\\"id\\\":.*\\}\\]", b.String(), "JSON output didn't match regular expression")
}

func TestPrintErrorJSON(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)

	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	f.PrintError("testing errors", fmt.Errorf("this is a test error %s", "TEST"))
	mockOut.Flush()

	assert.Regexp("^\\{\\\"type\\\":\\\"Error\\\",\\\"context\\\":\\\"testing errors\\\",\\\"message\\\":\\\"this is a test error TEST\\\"\\}", b.String(), "JSON output didn't match regular expression")
}

func TestPrintItemWrongBytesJSON(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintItem(make(chan int))
	assert.Error(err, "Should have gotten an error marshaling a JSON")
	mockOut.Flush()
}

func TestPrintListWrongBytesJSON(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(make(chan int))
	assert.Error(err, "Should have gotten an error marshaling a JSON")
	mockOut.Flush()
}

func TestPrintFatalJSON(t *testing.T) {

	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	osExit = func(code int) {
		got = code
	}
	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	f.PrintFatal("testing fatal", fmt.Errorf("this is a test error %s", "TEST"))
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}
