// +build windows

package brownfield

import (
	fw "github.com/ingrammicro/concerto/firewall"
)

func apply(p *fw.Policy) error {
	return p.Apply()
}
