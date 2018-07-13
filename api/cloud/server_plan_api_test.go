package cloud

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServerPlanServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServerPlanService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetServerPlanList(t *testing.T) {
	serverPlansIn := testdata.GetServerPlanData()
	for _, serverPlanIn := range *serverPlansIn {
		GetServerPlanListMocked(t, serverPlansIn, serverPlanIn.CloudProviderID)
		GetServerPlanListFailErrMocked(t, serverPlansIn, serverPlanIn.CloudProviderID)
		GetServerPlanListFailStatusMocked(t, serverPlansIn, serverPlanIn.CloudProviderID)
		GetServerPlanListFailJSONMocked(t, serverPlansIn, serverPlanIn.CloudProviderID)
	}
}

func TestGetServerPlan(t *testing.T) {
	serverPlansIn := testdata.GetServerPlanData()
	for _, serverPlanIn := range *serverPlansIn {
		GetServerPlanMocked(t, &serverPlanIn)
		GetServerPlanFailErrMocked(t, &serverPlanIn)
		GetServerPlanFailStatusMocked(t, &serverPlanIn)
		GetServerPlanFailJSONMocked(t, &serverPlanIn)
	}
}
