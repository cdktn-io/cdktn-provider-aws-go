// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafregionalrulegroup


type WafregionalRuleGroupActivatedRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafregional_rule_group#action WafregionalRuleGroup#action}
	Action *WafregionalRuleGroupActivatedRuleAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafregional_rule_group#priority WafregionalRuleGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafregional_rule_group#rule_id WafregionalRuleGroup#rule_id}.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafregional_rule_group#type WafregionalRuleGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

