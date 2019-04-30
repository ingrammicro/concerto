package firewall

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetPolicyMocked test mocked function
func GetPolicyMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// to json
	dIn, err := json.Marshal(policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Get", "/cloud/firewall_profile").Return(dIn, 200, nil)
	policyOut, err := ds.GetPolicy()
	policyIn.Md5 = policyOut.Md5
	assert.Nil(err, "Error getting firewall policy")
	assert.Equal(*policyIn, *policyOut, "GetPolicy returned different policies")

	return policyOut
}

// GetPolicyFailErrMocked test mocked function
func GetPolicyFailErrMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// to json
	dIn, err := json.Marshal(policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Get", "/cloud/firewall_profile").Return(dIn, 200, fmt.Errorf("mocked error"))
	policyOut, err := ds.GetPolicy()
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyOut
}

// GetPolicyFailStatusMocked test mocked function
func GetPolicyFailStatusMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// to json
	dIn, err := json.Marshal(policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Get", "/cloud/firewall_profile").Return(dIn, 499, nil)
	policyOut, err := ds.GetPolicy()
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyOut
}

// GetPolicyFailJSONMocked test mocked function
func GetPolicyFailJSONMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/cloud/firewall_profile").Return(dIn, 200, nil)
	policyOut, err := ds.GetPolicy()
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyOut
}

// AddPolicyRuleMocked test mocked function
func AddPolicyRuleMocked(t *testing.T, policyRuleIn *types.PolicyRule) *types.PolicyRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// to json
	dOut, err := json.Marshal(policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Post", "/cloud/firewall_profile/rules", mapIn).Return(dOut, 200, nil)
	policyRuleOut, err := ds.AddPolicyRule(mapIn)
	assert.Nil(err, "Error adding policy rule")
	assert.Equal(policyRuleIn, policyRuleOut, "AddPolicyRule returned different rules")

	return policyRuleOut
}

// AddPolicyRuleFailErrMocked test mocked function
func AddPolicyRuleFailErrMocked(t *testing.T, policyRuleIn *types.PolicyRule) *types.PolicyRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// to json
	dOut, err := json.Marshal(policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Post", "/cloud/firewall_profile/rules", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	policyRuleOut, err := ds.AddPolicyRule(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyRuleOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyRuleOut
}

// AddPolicyRuleFailStatusMocked test mocked function
func AddPolicyRuleFailStatusMocked(t *testing.T, policyRuleIn *types.PolicyRule) *types.PolicyRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// to json
	dOut, err := json.Marshal(policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Post", "/cloud/firewall_profile/rules", mapIn).Return(dOut, 499, nil)
	policyRuleOut, err := ds.AddPolicyRule(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyRuleOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyRuleOut
}

// AddPolicyRuleFailJSONMocked test mocked function
func AddPolicyRuleFailJSONMocked(t *testing.T, policyRuleIn *types.PolicyRule) *types.PolicyRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyRuleIn)
	assert.Nil(err, "Firewall test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/cloud/firewall_profile/rules", mapIn).Return(dIn, 200, nil)
	policyRuleOut, err := ds.AddPolicyRule(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyRuleOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyRuleOut
}

// UpdatePolicyMocked test mocked function
func UpdatePolicyMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// to json
	dOut, err := json.Marshal(policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Put", "/cloud/firewall_profile", mapIn).Return(dOut, 200, nil)
	policyOut, err := ds.UpdatePolicy(mapIn)
	assert.Nil(err, "Error updating policy")
	assert.Equal(policyIn, policyOut, "UpdatePolicy returned different policies")

	return policyOut
}

// UpdatePolicyFailErrMocked test mocked function
func UpdatePolicyFailErrMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// to json
	dOut, err := json.Marshal(policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Put", "/cloud/firewall_profile", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	policyOut, err := ds.UpdatePolicy(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyOut
}

// UpdatePolicyFailStatusMocked test mocked function
func UpdatePolicyFailStatusMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// to json
	dOut, err := json.Marshal(policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// call service
	cs.On("Put", "/cloud/firewall_profile", mapIn).Return(dOut, 499, nil)
	policyOut, err := ds.UpdatePolicy(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyOut
}

// UpdatePolicyFailJSONMocked test mocked function
func UpdatePolicyFailJSONMocked(t *testing.T, policyIn *types.Policy) *types.Policy {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallService(cs)
	assert.Nil(err, "Couldn't load firewall service")
	assert.NotNil(ds, "Firewall service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyIn)
	assert.Nil(err, "Firewall test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", "/cloud/firewall_profile", mapIn).Return(dIn, 200, nil)
	policyOut, err := ds.UpdatePolicy(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyOut
}
