// +build windows

package brownfield

import (
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/firewall"
)

func Apply(p *types.Policy) error {
	if len(p.Rules) > 0 {
		return firewall.Apply(*p)
	}
	return nil
}
