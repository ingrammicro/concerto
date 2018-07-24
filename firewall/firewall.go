package firewall

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/webservice"
)

const endpoint = "cloud/firewall_profile"

// Profile stores Firewall Profile data
type Profile struct {
	Profile Policy `json:"firewall_profile"`
}

// Policy stores Firewall Policy data
type Policy struct {
	Rules       []Rule `json:"rules"`
	Md5         string `json:"md5,omitempty"`
	ActualRules []Rule `json:"actual_rules,omitempty"`
}

// Rule stores Firewall Rule data
type Rule struct {
	Name     string `json:"name,omitempty"`
	Protocol string `json:"ip_protocol"`
	Cidr     string `json:"cidr_ip"`
	MinPort  int    `json:"min_port"`
	MaxPort  int    `json:"max_port"`
}

// Apply sets firewall policy
func (policy Policy) Apply() error {
	if len(policy.Rules) > 0 {
		return apply(policy)
	}
	return nil
}

// Flush purges firewall
func Flush() error {
	return flush()
}

func list(policy Policy) error {
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "CIDR\tPROTOCOL\tMIN\tMAX")

	for _, rule := range policy.ActualRules {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
	}
	w.Flush()
	return nil
}

func get() Policy {
	var policy Policy
	webservice, err := webservice.NewWebService()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Current firewall driver %s", driverName())
	data, _, err := webservice.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &policy)
	if err != nil {
		log.Fatal(err)
	}
	policy.Md5 = fmt.Sprintf("%x", md5.Sum(data))
	return policy
}

func cmdList(c *cli.Context) error {
	list(get())
	return nil
}

func cmdApply(c *cli.Context) error {
	policy := get()
	// Only apply firewall if we get a non-empty set of rules
	if len(policy.Rules) > 0 {
		apply(policy)
	} else {
		flush()
	}
	return nil
}

func cmdFlush(c *cli.Context) error {
	flush()
	return nil
}

func check(policy Policy, rule Rule) bool {
	exists := false
	for _, policyRule := range policy.Rules {
		if (policyRule.Cidr == rule.Cidr) && (policyRule.MaxPort == rule.MaxPort) && (policyRule.MinPort == rule.MinPort) && (policyRule.Protocol == rule.Protocol) {
			exists = true
		}
	}
	return exists
}

func cmdCheck(c *cli.Context) error {
	utils.FlagsRequired(c, []string{"cidr", "minPort", "maxPort", "ipProtocol"})

	newRule := &Rule{
		Protocol: c.String("ipProtocol"),
		Cidr:     c.String("cidr"),
		MinPort:  c.Int("minPort"),
		MaxPort:  c.Int("maxPort"),
	}
	policy := get()

	fmt.Printf("%t\n", check(policy, *newRule))
	return nil
}

func cmdAdd(c *cli.Context) error {
	utils.FlagsRequired(c, []string{"cidr", "minPort", "maxPort", "ipProtocol"})

	// API accepts only 1 rule
	newRule := &Rule{
		Protocol: c.String("ipProtocol"),
		Cidr:     c.String("cidr"),
		MinPort:  c.Int("minPort"),
		MaxPort:  c.Int("maxPort"),
	}
	policy := get()

	exists := check(policy, *newRule)

	if exists == false {
		policy.Rules = append(policy.Rules, *newRule)
		webservice, err := webservice.NewWebService()
		utils.CheckError(err)

		nRule := make(map[string]Rule)
		nRule["rule"] = *newRule

		json, err := json.Marshal(nRule)
		utils.CheckError(err)
		res, code, err := webservice.Post(fmt.Sprintf("%s/rules", endpoint), json)
		if res == nil {
			log.Fatal(err)
		}
		utils.CheckError(err)
		utils.CheckReturnCode(code, res)
	}

	return nil
}

func cmdUpdate(c *cli.Context) error {
	utils.FlagsRequired(c, []string{"rules"})

	fp := &Profile{
		Policy{},
	}

	var rules []Rule
	err := json.Unmarshal([]byte(c.String("rules")), &rules)
	utils.CheckError(err)
	fp.Profile.Rules = rules

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	json, err := json.Marshal(fp)
	utils.CheckError(err)
	res, code, err := webservice.Put(endpoint, json)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)
	return nil
}

func cmdRemove(c *cli.Context) error {
	utils.FlagsRequired(c, []string{"cidr", "minPort", "maxPort", "ipProtocol"})

	existingRule := &Rule{
		Protocol: c.String("ipProtocol"),
		Cidr:     c.String("cidr"),
		MinPort:  c.Int("minPort"),
		MaxPort:  c.Int("maxPort"),
	}
	policy := get()

	exists := check(policy, *existingRule)

	if exists == true {
		for i, rule := range policy.Rules {
			if rule == *existingRule {
				policy.Rules = append(policy.Rules[:i], policy.Rules[1+i:]...)
				break
			}
		}

		webservice, err := webservice.NewWebService()
		utils.CheckError(err)

		profile := &Profile{
			policy,
		}

		json, err := json.Marshal(profile)
		utils.CheckError(err)
		res, code, err := webservice.Put(endpoint, json)
		if res == nil {
			log.Fatal(err)
		}
		utils.CheckError(err)
		utils.CheckReturnCode(code, res)
	}
	return nil
}

// SubCommands return Firewall subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "apply",
			Usage:  "Applies selected firewall rules in host",
			Action: cmdApply,
		},
		{
			Name:   "flush",
			Usage:  "Flushes all firewall rules from host",
			Action: cmdFlush,
		},
		{
			Name:   "check",
			Usage:  "Checks if a firewall rule exists in host",
			Action: cmdCheck,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "minPort",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "maxPort",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ipProtocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "add",
			Usage:  "Adds a single firewall rule to host",
			Action: cmdAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "minPort",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "maxPort",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ipProtocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates all firewall rules",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "rules",
					Usage: `JSON array in the form '[{"ip_protocol":"...", "min_port":..., "max_port":..., "cidr_ip":"..."}, ... ]'`,
				},
			},
		},
		{
			Name:   "remove",
			Usage:  "Removes a firewall rule to host",
			Action: cmdRemove,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "minPort",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "maxPort",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ipProtocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "list",
			Usage:  "Lists all firewall rules associated to host",
			Action: cmdList,
		},
	}
}
