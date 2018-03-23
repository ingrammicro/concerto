// +build linux

package brownfield

import (
	fw "github.com/ingrammicro/concerto/firewall"
	"github.com/ingrammicro/concerto/utils"
)

func apply(p *fw.Policy) error {
	utils.RunCmd("/sbin/iptables -w -F INPUT")
	return p.Apply()
}
