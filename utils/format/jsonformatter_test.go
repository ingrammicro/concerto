package format

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/api/dns"
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPrintItemDomainJSON(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {

		domainOut := dns.GetDomainMocked(t, &domainIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("json", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*domainOut)
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

func TestPrintListDomainsJSON(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	domainOut := dns.GetDomainListMocked(t, domainsIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(domainOut)
	assert.Nil(err, "JSON formatter PrintItem error")
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
	assert.Nil(err, "JSON formatter PrintItem error")
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

	err := f.PrintItem(make(chan int))
	assert.Error(err, "Should have gotten an error marshaling a JSON")
	mockOut.Flush()
}
