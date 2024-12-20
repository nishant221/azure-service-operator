// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ServersIPV6FirewallRule_Spec struct {
	// Name: Resource name.
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *IPv6ServerFirewallRuleProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ServersIPV6FirewallRule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule ServersIPV6FirewallRule_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rule *ServersIPV6FirewallRule_Spec) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/ipv6FirewallRules"
func (rule *ServersIPV6FirewallRule_Spec) GetType() string {
	return "Microsoft.Sql/servers/ipv6FirewallRules"
}

// The properties of an IPv6 server firewall rule.
type IPv6ServerFirewallRuleProperties struct {
	// EndIPv6Address: The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to
	// startIpAddress.
	EndIPv6Address *string `json:"endIPv6Address,omitempty"`

	// StartIPv6Address: The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address *string `json:"startIPv6Address,omitempty"`
}
