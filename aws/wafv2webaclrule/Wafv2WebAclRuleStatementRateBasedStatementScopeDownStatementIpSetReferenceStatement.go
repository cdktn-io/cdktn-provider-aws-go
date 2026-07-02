// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement struct {
	// ARN of the IP set to reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#arn Wafv2WebAclRuleA#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// ip_set_forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#ip_set_forwarded_ip_config Wafv2WebAclRuleA#ip_set_forwarded_ip_config}
	IpSetForwardedIpConfig interface{} `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

