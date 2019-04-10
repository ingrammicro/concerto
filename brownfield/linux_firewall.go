// +build linux darwin

package brownfield

import (
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/utils"
)

func Apply(p *types.Policy) error {
	utils.RunCmd("/sbin/iptables -w -F INPUT")

	if len(p.Rules) > 0 {
		return firewall.Apply(*p)
	}
	return nil
}
