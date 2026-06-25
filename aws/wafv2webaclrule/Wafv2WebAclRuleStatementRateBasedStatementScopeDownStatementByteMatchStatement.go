// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement struct {
	// Area within the portion of a web request that you want AWS WAF to search for SearchString.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#positional_constraint Wafv2WebAclRuleA#positional_constraint}
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// String value to search for within the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#search_string Wafv2WebAclRuleA#search_string}
	SearchString *string `field:"required" json:"searchString" yaml:"searchString"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#field_to_match Wafv2WebAclRuleA#field_to_match}
	FieldToMatch interface{} `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#text_transformation Wafv2WebAclRuleA#text_transformation}
	TextTransformation interface{} `field:"optional" json:"textTransformation" yaml:"textTransformation"`
}

