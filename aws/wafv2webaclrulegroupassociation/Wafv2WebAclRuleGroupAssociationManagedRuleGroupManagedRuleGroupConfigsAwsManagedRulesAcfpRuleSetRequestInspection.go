// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspection struct {
	// Payload type for inspection, either JSON or FORM_ENCODED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/wafv2_web_acl_rule_group_association#payload_type Wafv2WebAclRuleGroupAssociation#payload_type}
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// address_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/wafv2_web_acl_rule_group_association#address_fields Wafv2WebAclRuleGroupAssociation#address_fields}
	AddressFields interface{} `field:"optional" json:"addressFields" yaml:"addressFields"`
	// email_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/wafv2_web_acl_rule_group_association#email_field Wafv2WebAclRuleGroupAssociation#email_field}
	EmailField interface{} `field:"optional" json:"emailField" yaml:"emailField"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/wafv2_web_acl_rule_group_association#password_field Wafv2WebAclRuleGroupAssociation#password_field}
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// phone_number_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/wafv2_web_acl_rule_group_association#phone_number_fields Wafv2WebAclRuleGroupAssociation#phone_number_fields}
	PhoneNumberFields interface{} `field:"optional" json:"phoneNumberFields" yaml:"phoneNumberFields"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/wafv2_web_acl_rule_group_association#username_field Wafv2WebAclRuleGroupAssociation#username_field}
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

