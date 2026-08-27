// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#payload_type Wafv2WebAclRuleA#payload_type}.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#password_field Wafv2WebAclRuleA#password_field}
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#username_field Wafv2WebAclRuleA#username_field}
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

