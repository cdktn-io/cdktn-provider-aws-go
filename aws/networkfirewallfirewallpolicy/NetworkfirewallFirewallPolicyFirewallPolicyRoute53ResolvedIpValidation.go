// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyRoute53ResolvedIpValidation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/networkfirewall_firewall_policy#action NetworkfirewallFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/networkfirewall_firewall_policy#dns_expiry_exception_policy NetworkfirewallFirewallPolicy#dns_expiry_exception_policy}.
	DnsExpiryExceptionPolicy *string `field:"required" json:"dnsExpiryExceptionPolicy" yaml:"dnsExpiryExceptionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/networkfirewall_firewall_policy#dns_firewall_rule_group_id NetworkfirewallFirewallPolicy#dns_firewall_rule_group_id}.
	DnsFirewallRuleGroupId *string `field:"required" json:"dnsFirewallRuleGroupId" yaml:"dnsFirewallRuleGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/networkfirewall_firewall_policy#bypass_destination_validation_ip_ranges NetworkfirewallFirewallPolicy#bypass_destination_validation_ip_ranges}.
	BypassDestinationValidationIpRanges *[]*string `field:"optional" json:"bypassDestinationValidationIpRanges" yaml:"bypassDestinationValidationIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/networkfirewall_firewall_policy#source_ip_ranges NetworkfirewallFirewallPolicy#source_ip_ranges}.
	SourceIpRanges *[]*string `field:"optional" json:"sourceIpRanges" yaml:"sourceIpRanges"`
}

