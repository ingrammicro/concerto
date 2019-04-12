package firewall

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// FirewallService manages firewall operations
type FirewallService struct {
	concertoService utils.ConcertoService
}

// NewFirewallService returns a Concerto firewall service
func NewFirewallService(concertoService utils.ConcertoService) (*FirewallService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &FirewallService{
		concertoService: concertoService,
	}, nil
}

// GetPolicy returns server firewall policy
func (fs *FirewallService) GetPolicy() (policy *types.Policy, err error) {
	log.Debug("GetPolicy")

	data, status, err := fs.concertoService.Get("/cloud/firewall_profile")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &policy); err != nil {
		return nil, err
	}
	policy.Md5 = fmt.Sprintf("%x", md5.Sum(data))

	return policy, nil
}

// AddPolicyRule adds a new firewall policy rule
func (fs *FirewallService) AddPolicyRule(ruleVector *map[string]interface{}) (policyRule *types.PolicyRule, err error) {
	log.Debug("AddPolicyRule")

	data, status, err := fs.concertoService.Post("/cloud/firewall_profile/rules", ruleVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &policyRule); err != nil {
		return nil, err
	}
	return policyRule, nil
}

// UpdatePolicy update server firewall profile
func (fs *FirewallService) UpdatePolicy(policyVector *map[string]interface{}) (policy *types.Policy, err error) {
	log.Debug("UpdatePolicy")

	data, status, err := fs.concertoService.Put("/cloud/firewall_profile", policyVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &policy); err != nil {
		return nil, err
	}
	return policy, nil
}
