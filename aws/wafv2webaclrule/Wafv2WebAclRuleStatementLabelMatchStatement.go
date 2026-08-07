// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementLabelMatchStatement struct {
	// String to match against. Must be 1-1024 characters and match pattern ^[0-9A-Za-z_\-:]+$.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/wafv2_web_acl_rule#key Wafv2WebAclRuleA#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specify whether to match using the label name or just the namespace. Valid values: LABEL, NAMESPACE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/wafv2_web_acl_rule#scope Wafv2WebAclRuleA#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

