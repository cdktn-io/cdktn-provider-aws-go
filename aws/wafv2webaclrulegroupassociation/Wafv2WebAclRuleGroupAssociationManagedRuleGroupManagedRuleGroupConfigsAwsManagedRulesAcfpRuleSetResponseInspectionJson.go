// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionJson struct {
	// Strings that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#failure_values Wafv2WebAclRuleGroupAssociation#failure_values}
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// Identifier of the JSON field to inspect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#identifier Wafv2WebAclRuleGroupAssociation#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Strings that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#success_values Wafv2WebAclRuleGroupAssociation#success_values}
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

