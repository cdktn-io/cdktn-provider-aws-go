// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCode struct {
	// Status codes that indicate a failed login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/wafv2_web_acl_rule_group_association#failure_codes Wafv2WebAclRuleGroupAssociation#failure_codes}
	FailureCodes *[]*float64 `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Status codes that indicate a successful login or account creation attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/wafv2_web_acl_rule_group_association#success_codes Wafv2WebAclRuleGroupAssociation#success_codes}
	SuccessCodes *[]*float64 `field:"required" json:"successCodes" yaml:"successCodes"`
}

