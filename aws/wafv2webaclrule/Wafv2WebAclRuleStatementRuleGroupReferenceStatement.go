// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRuleGroupReferenceStatement struct {
	// ARN of the RuleGroup (20-2048 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#arn Wafv2WebAclRuleA#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// excluded_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#excluded_rule Wafv2WebAclRuleA#excluded_rule}
	ExcludedRule interface{} `field:"optional" json:"excludedRule" yaml:"excludedRule"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#rule_action_override Wafv2WebAclRuleA#rule_action_override}
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
}

