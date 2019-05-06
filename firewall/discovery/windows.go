// +build windows

package discovery

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	windowsFirewallRuleNameHeader = "\r\nRule Name:"
)

var whiteSpaces = regexp.MustCompile("\\A\\s*\\z")
var ruleField = regexp.MustCompile("(?P<field>[^:]+):(?P<value>.*)")
var portRange = regexp.MustCompile("\\A(?P<min>[0-9]+)(-(?P<max>[0-9]+))?\\z")
var profileName = regexp.MustCompile(" Profile Settings:(\\z| )")
var profileState = regexp.MustCompile("\\AState")

func CurrentFirewallRules() ([]*FirewallChain, error) {
	profiles, err := enabledProfiles()
	if err != nil {
		return nil, err
	}
	fmt.Printf("DEBUG: discovered following enabled profiles %v\n", profiles)
	cmd := exec.Command("netsh", "advfirewall", "firewall", "show", "rule", "name=all", "dir=in")
	out := &bytes.Buffer{}
	cmd.Stdout = out
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("running netsh advfirewall firewall show rule command to obtain current rules: %v", err)
	}
	fwc := &FirewallChain{Name: "INPUT", Policy: "DROP"}
	for _, ruleText := range strings.Split(out.String(), windowsFirewallRuleNameHeader) {
		r, err := parseRule(ruleText, profiles)
		if err != nil {
			return nil, fmt.Errorf("parsing current firewall rule: %v", err)
		}
		if r != nil {
			fwc.Rules = append(fwc.Rules, r...)
		}
	}
	fmt.Printf("DEBUG: discovered %d rules\n", len(fwc.Rules))
	return []*FirewallChain{fwc}, nil
}

func enabledProfiles() ([]string, error) {
	cmd := exec.Command("netsh", "advfirewall", "show", "allprofiles")
	out := &bytes.Buffer{}
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("running netsh advfirewall show allprofiles command to obtain enabled profiles: %v", err)
	}
	var profiles []string
	var currentProfile string
	for _, l := range strings.Split(out.String(), "\r\n") {
		fmt.Printf("DEBUG: parsing %q\n", l)
		if profileName.MatchString(l) {
			currentProfile = strings.TrimSpace(strings.SplitN(l, " ", 2)[0])
			fmt.Printf("DEBUG: found profile name: %q\n", currentProfile)
			continue
		}
		if profileState.MatchString(l) {
			if currentProfile == "" {
				return nil, fmt.Errorf("could not parse netsh advfirewall show allprofiles command output: found state before rule name")
			}
			if strings.Contains(l, "ON") {
				profiles = append(profiles, currentProfile)
			}
		}
	}
	return profiles, nil
}

func parseRule(s string, profiles []string) ([]*FirewallRule, error) {
	if whiteSpaces.Match([]byte(s)) {
		return nil, nil
	}
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) < 3 {
		return nil, fmt.Errorf("cannot parse rule %q", s)
	}
	name := strings.TrimSpace(lines[0])
	ruleData := make(map[string]string)
	for _, line := range lines {
		match := ruleField.FindStringSubmatch(line)
		if match != nil {
			ruleData[strings.TrimSpace(match[1])] = strings.TrimSpace(match[2])
		}
	}
	var matchingProfile bool
	for _, p := range profiles {
		if strings.Contains(ruleData["Profiles"], p) {
			matchingProfile = true
			break
		}
	}
	if !matchingProfile {
		fmt.Printf("DEBUG: rule not belonging to enabled profile (rule's %q vs enabled %q)\n", profiles, ruleData["Profiles"])
		return nil, nil
	}
	if v := ruleData["Enabled"]; v != "Yes" {
		return nil, nil
	}
	protocol := ruleData["Protocol"]
	if protocol != "TCP" && protocol != "UDP" {
		return nil, nil
	}
	localPort := ruleData["LocalPort"]
	if localPort == "" || localPort == "Any" {
		localPort = "1-65535"
	}
	var rules []*FirewallRule
	for _, port := range strings.Split(localPort, ",") {
		match := portRange.FindStringSubmatch(port)
		if match == nil {
			// If rule contains dynamic port (such as RPC), do not consider any part of it
			//fmt.Printf("Encountered dynamic port rule:\n%s\n", s)
			return nil, nil
		}
		minPort, _ := strconv.Atoi(match[1])
		maxPort := minPort
		if match[3] != "" {
			maxPort, _ = strconv.Atoi(match[3])
		}
		cidr := ruleData["RemoteIp"]
		if cidr == "" || cidr == "Any" {
			cidr = "0.0.0.0/0"
		}
		r := &FirewallRule{
			Name:     name,
			Target:   "ACCEPT",
			Protocol: protocol,
			Source:   cidr,
			Dports:   [2]int{minPort, maxPort},
		}
		fmt.Printf("DEBUG: Parsed rule: %v\n", *r)
		rules = append(rules, r)
	}
	return rules, nil
}
