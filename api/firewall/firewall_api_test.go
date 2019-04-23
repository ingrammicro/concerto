package firewall

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDispatcherServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewFirewallService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetPolicy(t *testing.T) {
	pIn := testdata.GetPolicyData()
	GetPolicyMocked(t, pIn)
	GetPolicyFailErrMocked(t, pIn)
	GetPolicyFailStatusMocked(t, pIn)
	GetPolicyFailJSONMocked(t, pIn)
}

func TestAddPolicyRule(t *testing.T) {
	pIn := testdata.GetPolicyData()
	for _, pr := range pIn.Rules {
		AddPolicyRuleMocked(t, &pr)
		AddPolicyRuleFailErrMocked(t, &pr)
		AddPolicyRuleFailStatusMocked(t, &pr)
		AddPolicyRuleFailJSONMocked(t, &pr)
	}
}

func TestUpdatePolicy(t *testing.T) {
	pIn := testdata.GetPolicyData()
	UpdatePolicyMocked(t, pIn)
	UpdatePolicyFailErrMocked(t, pIn)
	UpdatePolicyFailStatusMocked(t, pIn)
	UpdatePolicyFailJSONMocked(t, pIn)
}
