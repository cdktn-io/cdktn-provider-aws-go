// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatement struct {
	// ARN of the RegexPatternSet (20-2048 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#arn Wafv2WebAclRuleA#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#field_to_match Wafv2WebAclRuleA#field_to_match}
	FieldToMatch interface{} `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#text_transformation Wafv2WebAclRuleA#text_transformation}
	TextTransformation interface{} `field:"optional" json:"textTransformation" yaml:"textTransformation"`
}

