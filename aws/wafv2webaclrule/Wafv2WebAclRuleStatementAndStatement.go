// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementAndStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule#statement Wafv2WebAclRuleA#statement}
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

