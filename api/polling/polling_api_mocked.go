package polling

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// PingMocked test mocked function
func PingMocked(t *testing.T, pingIn *types.PollingPing) *types.PollingPing {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dOut, err := json.Marshal(pingIn)
	assert.Nil(err, "Polling test data corrupted")

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", "/command_polling/pings", &payload).Return(dOut, 201, nil)
	pingOut, status, err := ds.Ping()
	assert.Nil(err, "Error getting ping")
	assert.Equal(status, 201, "Ping returned invalid response")
	assert.Equal(pingIn.PendingCommands, true, "Ping returned no pending command available")
	pingIn.PendingCommands = false
	assert.Equal(pingIn.PendingCommands, false, "Ping returned pending command available")

	return pingOut
}

// PingFailErrMocked test mocked function
func PingFailErrMocked(t *testing.T, pingIn *types.PollingPing) *types.PollingPing {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(pingIn)
	assert.Nil(err, "Polling test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", "/command_polling/pings", &payload).Return(dIn, 404, fmt.Errorf("Mocked error"))
	pingOut, _, err := ds.Ping()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(pingOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return pingOut
}

// PingFailStatusMocked test mocked function
func PingFailStatusMocked(t *testing.T, pingIn *types.PollingPing) *types.PollingPing {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(pingIn)
	assert.Nil(err, "Polling test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", "/command_polling/pings", &payload).Return(dIn, 499, fmt.Errorf("Error 499 Mocked error"))
	pingOut, status, err := ds.Ping()

	assert.Equal(status, 499, "Ping returned an unexpected status code")
	assert.NotNil(err, "We are expecting a status code error")
	assert.Nil(pingOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return pingOut
}

// PingFailJSONMocked test mocked function
func PingFailJSONMocked(t *testing.T, pingIn *types.PollingPing) *types.PollingPing {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", "/command_polling/pings", &payload).Return(dIn, 201, nil)
	pingOut, _, err := ds.Ping()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(pingOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return pingOut
}

// GetNextCommandMocked test mocked function
func GetNextCommandMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dOut, err := json.Marshal(commandIn)
	assert.Nil(err, "GetNextCommand test data corrupted")

	// call service
	cs.On("Get", "/command_polling/command").Return(dOut, 200, nil)
	commandOut, status, err := ds.GetNextCommand()
	assert.Nil(err, "Error getting polling command")
	assert.Equal(status, 200, "GetNextCommand returned invalid response")
	assert.Equal(*commandIn, *commandOut, "GetNextCommand returned different nodes")

	return commandOut
}

// GetNextCommandFailErrMocked test mocked function
func GetNextCommandFailErrMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "GetNextCommand test data corrupted")

	dIn = nil

	// call service
	cs.On("Get", "/command_polling/command").Return(dIn, 404, fmt.Errorf("Mocked error"))
	commandOut, _, err := ds.GetNextCommand()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return commandOut
}

// GetNextCommandFailStatusMocked test mocked function
func GetNextCommandFailStatusMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "GetNextCommand test data corrupted")

	dIn = nil

	// call service
	cs.On("Get", "/command_polling/command").Return(dIn, 499, fmt.Errorf("Error 499 Mocked error"))
	commandOut, status, err := ds.GetNextCommand()

	assert.Equal(status, 499, "GetNextCommand returned an unexpected status code")
	assert.NotNil(err, "We are expecting a status code error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return commandOut
}

// GetNextCommandFailJSONMocked test mocked function
func GetNextCommandFailJSONMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/command_polling/command").Return(dIn, 200, nil)
	commandOut, _, err := ds.GetNextCommand()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return commandOut
}

// UpdateCommandMocked test mocked function
func UpdateCommandMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dOut, err := json.Marshal(commandIn)
	assert.Nil(err, "UpdateCommand test data corrupted")

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/command_polling/commands/%s", commandIn.ID), &payload).Return(dOut, 200, nil)
	commandOut, status, err := ds.UpdateCommand(&payload, commandIn.ID)
	assert.Nil(err, "Error getting polling command")
	assert.Equal(status, 200, "UpdateCommand returned invalid response")
	assert.Equal(*commandIn, *commandOut, "UpdateCommand returned different nodes")

	return commandOut
}

// UpdateCommandFailErrMocked test mocked function
func UpdateCommandFailErrMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "UpdateCommand test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/command_polling/commands/%s", commandIn.ID), &payload).Return(dIn, 400, fmt.Errorf("Mocked error"))
	commandOut, _, err := ds.UpdateCommand(&payload, commandIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return commandOut
}

// UpdateCommandFailStatusMocked test mocked function
func UpdateCommandFailStatusMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "UpdateCommand test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/command_polling/commands/%s", commandIn.ID), &payload).Return(dIn, 499, fmt.Errorf("Error 499 Mocked error"))
	commandOut, status, err := ds.UpdateCommand(&payload, commandIn.ID)

	assert.Equal(status, 499, "UpdateCommand returned an unexpected status code")
	assert.NotNil(err, "We are expecting a status code error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return commandOut
}

// UpdateCommandFailJSONMocked test mocked function
func UpdateCommandFailJSONMocked(t *testing.T, commandIn *types.PollingCommand) *types.PollingCommand {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	payload := make(map[string]interface{})
	cs.On("Put", fmt.Sprintf("/command_polling/commands/%s", commandIn.ID), &payload).Return(dIn, 200, nil)
	commandOut, _, err := ds.UpdateCommand(&payload, commandIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return commandOut
}

// ReportBootstrapLogMocked test mocked function
func ReportBootstrapLogMocked(t *testing.T, commandIn *types.PollingContinuousReport) *types.PollingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dOut, err := json.Marshal(commandIn)
	assert.Nil(err, "ReportBootstrapLog test data corrupted")

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/command_polling/bootstrap_logs"), &payload).Return(dOut, 201, nil)
	commandOut, status, err := ds.ReportBootstrapLog(&payload)

	assert.Nil(err, "Error posting report command")
	assert.Equal(status, 201, "ReportBootstrapLog returned invalid response")
	assert.Equal(commandOut.Stdout, "Bootstrap log created", "ReportBootstrapLog returned unexpected message")

	return commandOut
}

// ReportBootstrapLogFailErrMocked test mocked function
func ReportBootstrapLogFailErrMocked(t *testing.T, commandIn *types.PollingContinuousReport) *types.PollingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "ReportBootstrapLog test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/command_polling/bootstrap_logs"), &payload).Return(dIn, 400, fmt.Errorf("Mocked error"))
	commandOut, _, err := ds.ReportBootstrapLog(&payload)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return commandOut
}

// ReportBootstrapLogFailStatusMocked test mocked function
func ReportBootstrapLogFailStatusMocked(t *testing.T, commandIn *types.PollingContinuousReport) *types.PollingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// to json
	dIn, err := json.Marshal(commandIn)
	assert.Nil(err, "ReportBootstrapLog test data corrupted")

	dIn = nil

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/command_polling/bootstrap_logs"), &payload).Return(dIn, 499, fmt.Errorf("Error 499 Mocked error"))
	commandOut, status, err := ds.ReportBootstrapLog(&payload)

	assert.Equal(status, 499, "ReportBootstrapLog returned an unexpected status code")
	assert.NotNil(err, "We are expecting a status code error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return commandOut
}

// ReportBootstrapLogFailJSONMocked test mocked function
func ReportBootstrapLogFailJSONMocked(t *testing.T, commandIn *types.PollingContinuousReport) *types.PollingContinuousReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPollingService(cs)
	assert.Nil(err, "Couldn't load polling service")
	assert.NotNil(ds, "Polling service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	payload := make(map[string]interface{})
	cs.On("Post", fmt.Sprintf("/command_polling/bootstrap_logs"), &payload).Return(dIn, 201, nil)
	commandOut, _, err := ds.ReportBootstrapLog(&payload)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(commandOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return commandOut
}
