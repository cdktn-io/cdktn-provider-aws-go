// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatement struct {
	// and_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#and_statement Wafv2WebAclRuleA#and_statement}
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// asn_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#asn_match_statement Wafv2WebAclRuleA#asn_match_statement}
	AsnMatchStatement interface{} `field:"optional" json:"asnMatchStatement" yaml:"asnMatchStatement"`
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#byte_match_statement Wafv2WebAclRuleA#byte_match_statement}
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#geo_match_statement Wafv2WebAclRuleA#geo_match_statement}
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#ip_set_reference_statement Wafv2WebAclRuleA#ip_set_reference_statement}
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#label_match_statement Wafv2WebAclRuleA#label_match_statement}
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// managed_rule_group_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#managed_rule_group_statement Wafv2WebAclRuleA#managed_rule_group_statement}
	ManagedRuleGroupStatement interface{} `field:"optional" json:"managedRuleGroupStatement" yaml:"managedRuleGroupStatement"`
	// not_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#not_statement Wafv2WebAclRuleA#not_statement}
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// or_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#or_statement Wafv2WebAclRuleA#or_statement}
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
	// rate_based_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#rate_based_statement Wafv2WebAclRuleA#rate_based_statement}
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#regex_match_statement Wafv2WebAclRuleA#regex_match_statement}
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#regex_pattern_set_reference_statement Wafv2WebAclRuleA#regex_pattern_set_reference_statement}
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// rule_group_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#rule_group_reference_statement Wafv2WebAclRuleA#rule_group_reference_statement}
	RuleGroupReferenceStatement interface{} `field:"optional" json:"ruleGroupReferenceStatement" yaml:"ruleGroupReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#size_constraint_statement Wafv2WebAclRuleA#size_constraint_statement}
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#sqli_match_statement Wafv2WebAclRuleA#sqli_match_statement}
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule#xss_match_statement Wafv2WebAclRuleA#xss_match_statement}
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

