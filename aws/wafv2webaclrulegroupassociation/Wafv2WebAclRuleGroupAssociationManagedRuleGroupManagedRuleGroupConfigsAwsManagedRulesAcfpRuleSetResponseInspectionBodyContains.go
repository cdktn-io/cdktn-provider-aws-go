// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionBodyContains struct {
	// Strings that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/wafv2_web_acl_rule_group_association#failure_strings Wafv2WebAclRuleGroupAssociation#failure_strings}
	FailureStrings *[]*string `field:"required" json:"failureStrings" yaml:"failureStrings"`
	// Strings that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/wafv2_web_acl_rule_group_association#success_strings Wafv2WebAclRuleGroupAssociation#success_strings}
	SuccessStrings *[]*string `field:"required" json:"successStrings" yaml:"successStrings"`
}

