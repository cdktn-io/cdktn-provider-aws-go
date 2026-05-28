// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverride struct {
	// Name of the rule to override (1-128 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/wafv2_web_acl_rule#name Wafv2WebAclRuleA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// action_to_use block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/wafv2_web_acl_rule#action_to_use Wafv2WebAclRuleA#action_to_use}
	ActionToUse interface{} `field:"optional" json:"actionToUse" yaml:"actionToUse"`
}

