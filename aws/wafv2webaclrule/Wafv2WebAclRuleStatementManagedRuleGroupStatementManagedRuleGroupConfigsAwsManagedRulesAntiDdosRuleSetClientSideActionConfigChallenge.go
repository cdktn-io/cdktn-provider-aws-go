// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAntiDdosRuleSetClientSideActionConfigChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafv2_web_acl_rule#usage_of_action Wafv2WebAclRuleA#usage_of_action}.
	UsageOfAction *string `field:"required" json:"usageOfAction" yaml:"usageOfAction"`
	// exempt_uri_regular_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafv2_web_acl_rule#exempt_uri_regular_expression Wafv2WebAclRuleA#exempt_uri_regular_expression}
	ExemptUriRegularExpression interface{} `field:"optional" json:"exemptUriRegularExpression" yaml:"exemptUriRegularExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/wafv2_web_acl_rule#sensitivity Wafv2WebAclRuleA#sensitivity}.
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
}

