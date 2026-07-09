// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspection struct {
	// body_contains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule#body_contains Wafv2WebAclRuleA#body_contains}
	BodyContains interface{} `field:"optional" json:"bodyContains" yaml:"bodyContains"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule#header Wafv2WebAclRuleA#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule#json Wafv2WebAclRuleA#json}
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule#status_code Wafv2WebAclRuleA#status_code}
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

