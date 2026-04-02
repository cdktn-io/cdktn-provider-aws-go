// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/wafv2_web_acl_rule#failure_codes Wafv2WebAclRuleA#failure_codes}.
	FailureCodes *[]*float64 `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/wafv2_web_acl_rule#success_codes Wafv2WebAclRuleA#success_codes}.
	SuccessCodes *[]*float64 `field:"required" json:"successCodes" yaml:"successCodes"`
}

