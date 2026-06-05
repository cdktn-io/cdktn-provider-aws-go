// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRateBasedStatement struct {
	// Setting that indicates how to aggregate the request counts. Valid values: IP, FORWARDED_IP, CUSTOM_KEYS, CONSTANT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#aggregate_key_type Wafv2WebAclRuleA#aggregate_key_type}
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// Rate limit threshold (10-2000000000).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#limit Wafv2WebAclRuleA#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// custom_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#custom_keys Wafv2WebAclRuleA#custom_keys}
	CustomKeys interface{} `field:"optional" json:"customKeys" yaml:"customKeys"`
	// Time window for AWS WAF to use to check the rate (60, 120, 300, 600).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#evaluation_window_sec Wafv2WebAclRuleA#evaluation_window_sec}
	EvaluationWindowSec *float64 `field:"optional" json:"evaluationWindowSec" yaml:"evaluationWindowSec"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#forwarded_ip_config Wafv2WebAclRuleA#forwarded_ip_config}
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule#scope_down_statement Wafv2WebAclRuleA#scope_down_statement}
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

