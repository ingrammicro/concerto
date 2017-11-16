package polling

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPollingServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewPollingService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestPing(t *testing.T) {
	pingIn := testdata.GetPollingPingData()
	PingMocked(t, pingIn)
	PingFailErrMocked(t, pingIn)
	PingFailStatusMocked(t, pingIn)
	PingFailJSONMocked(t, pingIn)
}

func TestGetNextCommand(t *testing.T) {
	commandIn := testdata.GetPollingCommandData()
	GetNextCommandMocked(t, commandIn)
	GetNextCommandFailErrMocked(t, commandIn)
	GetNextCommandFailStatusMocked(t, commandIn)
	GetNextCommandFailJSONMocked(t, commandIn)
}

func TestUpdateCommand(t *testing.T) {
	commandIn := testdata.GetPollingCommandData()
	UpdateCommandMocked(t, commandIn)
	UpdateCommandFailErrMocked(t, commandIn)
	UpdateCommandFailStatusMocked(t, commandIn)
	UpdateCommandFailJSONMocked(t, commandIn)
}
