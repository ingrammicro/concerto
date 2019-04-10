// +build darwin

package firewall

import (
	"fmt"
	"github.com/ingrammicro/concerto/api/types"
)

func driverName() string {
	return "darwin"
}

func Apply(policy types.Policy) error {
	fmt.Println("iptables -A INPUT -i lo -j ACCEPT")
	fmt.Println("iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")

	for _, rule := range policy.Rules {
		fmt.Printf("iptables -A INPUT -s %s -p %s --dport %d:%d -j ACCEPT\n", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
	}
	fmt.Println("iptables -P INPUT ACCEPT")
	fmt.Println("iptables -F INPUT")
	return nil
}

func flush() error {
	fmt.Println("iptables -F INPUT")
	fmt.Println("iptables -P INPUT DROP")
	return nil
}
