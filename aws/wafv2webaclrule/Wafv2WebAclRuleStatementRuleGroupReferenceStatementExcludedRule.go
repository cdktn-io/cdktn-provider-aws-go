// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRuleGroupReferenceStatementExcludedRule struct {
	// Name of the rule to exclude (1-128 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/wafv2_web_acl_rule#name Wafv2WebAclRuleA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

