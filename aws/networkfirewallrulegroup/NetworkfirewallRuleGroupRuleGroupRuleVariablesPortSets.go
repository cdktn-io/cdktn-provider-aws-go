// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRuleVariablesPortSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/networkfirewall_rule_group#key NetworkfirewallRuleGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// port_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/networkfirewall_rule_group#port_set NetworkfirewallRuleGroup#port_set}
	PortSet *NetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsPortSet `field:"required" json:"portSet" yaml:"portSet"`
}

