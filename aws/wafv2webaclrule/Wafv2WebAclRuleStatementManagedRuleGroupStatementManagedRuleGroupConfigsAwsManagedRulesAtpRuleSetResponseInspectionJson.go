// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionJson struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/wafv2_web_acl_rule#failure_values Wafv2WebAclRuleA#failure_values}.
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/wafv2_web_acl_rule#identifier Wafv2WebAclRuleA#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/wafv2_web_acl_rule#success_values Wafv2WebAclRuleA#success_values}.
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

