// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementFieldToMatchJsonBody struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#match_scope Wafv2WebAclRuleA#match_scope}.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#invalid_fallback_behavior Wafv2WebAclRuleA#invalid_fallback_behavior}.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
	// match_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#match_pattern Wafv2WebAclRuleA#match_pattern}
	MatchPattern interface{} `field:"optional" json:"matchPattern" yaml:"matchPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#oversize_handling Wafv2WebAclRuleA#oversize_handling}.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

