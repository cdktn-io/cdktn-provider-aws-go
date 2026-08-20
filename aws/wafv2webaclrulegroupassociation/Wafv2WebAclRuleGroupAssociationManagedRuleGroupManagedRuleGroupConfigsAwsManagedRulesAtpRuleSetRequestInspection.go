// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspection struct {
	// Payload type for inspection, either JSON or FORM_ENCODED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/wafv2_web_acl_rule_group_association#payload_type Wafv2WebAclRuleGroupAssociation#payload_type}
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/wafv2_web_acl_rule_group_association#password_field Wafv2WebAclRuleGroupAssociation#password_field}
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/wafv2_web_acl_rule_group_association#username_field Wafv2WebAclRuleGroupAssociation#username_field}
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

