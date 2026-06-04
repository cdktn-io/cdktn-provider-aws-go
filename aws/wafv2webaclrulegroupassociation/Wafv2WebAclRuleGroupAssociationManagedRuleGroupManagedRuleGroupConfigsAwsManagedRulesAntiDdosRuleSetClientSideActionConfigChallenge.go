// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAntiDdosRuleSetClientSideActionConfigChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/wafv2_web_acl_rule_group_association#usage_of_action Wafv2WebAclRuleGroupAssociation#usage_of_action}.
	UsageOfAction *string `field:"required" json:"usageOfAction" yaml:"usageOfAction"`
	// exempt_uri_regular_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/wafv2_web_acl_rule_group_association#exempt_uri_regular_expression Wafv2WebAclRuleGroupAssociation#exempt_uri_regular_expression}
	ExemptUriRegularExpression interface{} `field:"optional" json:"exemptUriRegularExpression" yaml:"exemptUriRegularExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/wafv2_web_acl_rule_group_association#sensitivity Wafv2WebAclRuleGroupAssociation#sensitivity}.
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
}

