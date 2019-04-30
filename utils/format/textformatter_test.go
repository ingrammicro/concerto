package format

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPrintItemTXT(t *testing.T) {

	assert := assert.New(t)
	serversIn := testdata.GetServerData()
	for _, serverIn := range serversIn {

		serverOut := cloud.GetServerMocked(t, serverIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*serverOut)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^ID:\\ *.*\n*.\n", b.String(), "Text output didn't match regular expression")

	}
}

func TestPrintItemTemplateTXT(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range templatesIn {

		templateOut := blueprint.GetTemplateMocked(t, templateIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*templateOut)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^ID:\\ *.*\n*.\n", b.String(), "Text output didn't match regular expression")

	}
}

func TestPrintListTXT(t *testing.T) {

	assert := assert.New(t)
	serversIn := testdata.GetServerData()
	serversOut := cloud.GetServerListMocked(t, serversIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(serversOut)
	assert.Nil(err, "Text formatter PrintList error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp(fmt.Sprintf("^ID.*\n%s.*\n.*", serversOut[0].ID), b.String(), "Text output didn't match regular expression")
}

func TestPrintListTemplateTXT(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	templatesOut := blueprint.GetTemplateListMocked(t, templatesIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(templatesOut)
	assert.Nil(err, "Text formatter PrintList error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp(fmt.Sprintf("^ID.*\n%s.*\n.*", templatesOut[0].ID), b.String(), "Text output didn't match regular expression")
}

func TestPrintListTemplateScriptsTXT(t *testing.T) {

	assert := assert.New(t)
	tScriptsIn := testdata.GetTemplateScriptData()

	for _, tsIn := range tScriptsIn {
		tScriptsOut := blueprint.GetTemplateScriptListMocked(t, tScriptsIn, tsIn.ID, tsIn.Type)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintList(tScriptsOut)
		assert.Nil(err, "Text formatter PrintList error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp(fmt.Sprintf("^ID.*\n%s.*\n.*", tScriptsOut[0].ID), b.String(), "Text output didn't match regular expression")

	}
}

func TestPrintListNonSliceErrorTXT(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList("string")
	assert.Error(err, "A 'non slice' error should have arisen")
	mockOut.Flush()

}

func TestPrintError(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)

	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	f.PrintError("testing errors", fmt.Errorf("this is a test error %s", "TEST"))
	mockOut.Flush()

	assert.Regexp("^ERROR:.*\n -> .*\n", b.String(), "Text output didn't match regular expression")
}

func TestPrintListMinifySeconds(t *testing.T) {

	assert := assert.New(t)
	dummyData := testdata.GetDummyData()

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(dummyData)
	assert.Nil(err, "Text formatter PrintList error")
	mockOut.Flush()
}

func TestPrintItemJSONRawMessage(t *testing.T) {

	assert := assert.New(t)
	dummyData := testdata.GetDummyData()

	for _, dummy := range dummyData {
		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*dummy)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()
	}
}

func TestPrintListJSONRawMessage(t *testing.T) {

	assert := assert.New(t)
	dummyData := testdata.GetDummyData()

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(dummyData)
	assert.Nil(err, "Text formatter PrintList error")
	mockOut.Flush()
}

func TestPrintListJSONRawMessageNil(t *testing.T) {

	assert := assert.New(t)
	dummyData := testdata.GetDummyData()

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(dummyData)
	assert.Nil(err, "Text formatter PrintList error")
	mockOut.Flush()
}

func TestPrintFatalTXT(t *testing.T) {

	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	osExit = func(code int) {
		got = code
	}
	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	f.PrintFatal("testing fatal", fmt.Errorf("this is a test error %s", "TEST"))
	if exp := 1; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}
