package storage

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetVolumeListMocked test mocked function
func GetVolumeListMocked(t *testing.T, volumesIn []*types.Volume) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", "/storage/volumes").Return(dIn, 200, nil)
	volumesOut, err := ds.GetVolumeList("")
	assert.Nil(err, "Error getting volume list")
	assert.Equal(volumesIn, volumesOut, "GetVolumeList returned different volumes")

	return volumesOut
}

func GetVolumeListMockedFilteredByServer(t *testing.T, volumesIn []*types.Volume) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/servers/%s/volumes", volumesIn[0].AttachedServerID)).Return(dIn, 200, nil)
	volumesOut, err := ds.GetVolumeList(volumesIn[0].AttachedServerID)
	assert.Nil(err, "Error getting volume list filtered by server")
	assert.Equal(volumesIn, volumesOut, "GetVolumeList returned different volumes")

	return volumesOut
}

// GetVolumeListFailErrMocked test mocked function
func GetVolumeListFailErrMocked(t *testing.T, volumesIn []*types.Volume) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", "/storage/volumes").Return(dIn, 200, fmt.Errorf("mocked error"))
	volumesOut, err := ds.GetVolumeList("")

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(volumesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return volumesOut
}

// GetVolumeListFailStatusMocked test mocked function
func GetVolumeListFailStatusMocked(t *testing.T, volumesIn []*types.Volume) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumesIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", "/storage/volumes").Return(dIn, 499, nil)
	volumesOut, err := ds.GetVolumeList("")

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(volumesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return volumesOut
}

// GetVolumeListFailJSONMocked test mocked function
func GetVolumeListFailJSONMocked(t *testing.T, volumesIn []*types.Volume) []*types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/storage/volumes").Return(dIn, 200, nil)
	volumesOut, err := ds.GetVolumeList("")

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(volumesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return volumesOut
}

// GetVolumeMocked test mocked function
func GetVolumeMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 200, nil)
	volumeOut, err := ds.GetVolume(volumeIn.ID)
	assert.Nil(err, "Error getting volume")
	assert.Equal(*volumeIn, *volumeOut, "GetVolume returned different volumes")

	return volumeOut
}

// GetVolumeFailErrMocked test mocked function
func GetVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	volumeOut, err := ds.GetVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return volumeOut
}

// GetVolumeFailStatusMocked test mocked function
func GetVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 499, nil)
	volumeOut, err := ds.GetVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return volumeOut
}

// GetVolumeFailJSONMocked test mocked function
func GetVolumeFailJSONMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 200, nil)
	volumeOut, err := ds.GetVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return volumeOut
}

// CreateVolumeMocked test mocked function
func CreateVolumeMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", "/storage/volumes/", mapIn).Return(dOut, 200, nil)
	volumeOut, err := ds.CreateVolume(mapIn)
	assert.Nil(err, "Error creating volume list")
	assert.Equal(volumeIn, volumeOut, "CreateVolume returned different volumes")

	return volumeOut
}

// CreateVolumeFailErrMocked test mocked function
func CreateVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", "/storage/volumes/", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	volumeOut, err := ds.CreateVolume(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return volumeOut
}

// CreateVolumeFailStatusMocked test mocked function
func CreateVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", "/storage/volumes/", mapIn).Return(dOut, 499, nil)
	volumeOut, err := ds.CreateVolume(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return volumeOut
}

// CreateVolumeFailJSONMocked test mocked function
func CreateVolumeFailJSONMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/storage/volumes/", mapIn).Return(dIn, 200, nil)
	volumeOut, err := ds.CreateVolume(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return volumeOut
}

// UpdateVolumeMocked test mocked function
func UpdateVolumeMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID), mapIn).Return(dOut, 200, nil)
	volumeOut, err := ds.UpdateVolume(mapIn, volumeIn.ID)
	assert.Nil(err, "Error updating volume list")
	assert.Equal(volumeIn, volumeOut, "UpdateVolume returned different volumes")

	return volumeOut
}

// UpdateVolumeFailErrMocked test mocked function
func UpdateVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	volumeOut, err := ds.UpdateVolume(mapIn, volumeIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return volumeOut
}

// UpdateVolumeFailStatusMocked test mocked function
func UpdateVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID), mapIn).Return(dOut, 499, nil)
	volumeOut, err := ds.UpdateVolume(mapIn, volumeIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return volumeOut
}

// UpdateVolumeFailJSONMocked test mocked function
func UpdateVolumeFailJSONMocked(t *testing.T, volumeIn *types.Volume) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID), mapIn).Return(dIn, 200, nil)
	volumeOut, err := ds.UpdateVolume(mapIn, volumeIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return volumeOut
}

// AttachVolumeMocked test mocked function
func AttachVolumeMocked(t *testing.T, volumeIn *types.Volume) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: volumeIn.AttachedServerID})
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.AttachVolume(mapIn, volumeIn.ID)
	assert.Nil(err, "Error attaching volume")
	assert.Equal(volumeIn.AttachedServerID, serverOut.ID, "AttachVolume returned invalid values")

	return serverOut
}

// AttachVolumeFailErrMocked test mocked function
func AttachVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: volumeIn.AttachedServerID})
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.AttachVolume(mapIn, volumeIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// AttachVolumeFailStatusMocked test mocked function
func AttachVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: volumeIn.AttachedServerID})
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.AttachVolume(mapIn, volumeIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return serverOut
}

// AttachVolumeFailJSONMocked test mocked function
func AttachVolumeFailJSONMocked(t *testing.T, volumeIn *types.Volume) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.AttachVolume(mapIn, volumeIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// DetachVolumeMocked test mocked function
func DetachVolumeMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID)).Return(dIn, 200, nil)
	err = ds.DetachVolume(volumeIn.ID)
	assert.Nil(err, "Error detaching volume")
}

// DetachVolumeFailErrMocked test mocked function
func DetachVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DetachVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DetachVolumeFailStatusMocked test mocked function
func DetachVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s/attached_server", volumeIn.ID)).Return(dIn, 499, nil)
	err = ds.DetachVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteVolumeMocked test mocked function
func DeleteVolumeMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteVolume(volumeIn.ID)
	assert.Nil(err, "Error deleting volume")
}

// DeleteVolumeFailErrMocked test mocked function
func DeleteVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteVolumeFailStatusMocked test mocked function
func DeleteVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s", volumeIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DiscardVolumeMocked test mocked function
func DiscardVolumeMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s/discard", volumeIn.ID)).Return(dIn, 200, nil)
	err = ds.DiscardVolume(volumeIn.ID)
	assert.Nil(err, "Error discarding volume")
}

// DiscardVolumeFailErrMocked test mocked function
func DiscardVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s/discard", volumeIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DiscardVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DiscardVolumeFailStatusMocked test mocked function
func DiscardVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewVolumeService(cs)
	assert.Nil(err, "Couldn't load volume service")
	assert.NotNil(ds, "Volume service not instanced")

	// to json
	dIn, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/storage/volumes/%s/discard", volumeIn.ID)).Return(dIn, 499, nil)
	err = ds.DiscardVolume(volumeIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
