// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2ActionFindingFieldsUpdate struct {
	// A comment for the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#comment SecurityhubAutomationRuleV2#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The severity ID to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#severity_id SecurityhubAutomationRuleV2#severity_id}
	SeverityId *float64 `field:"optional" json:"severityId" yaml:"severityId"`
	// The status ID to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#status_id SecurityhubAutomationRuleV2#status_id}
	StatusId *float64 `field:"optional" json:"statusId" yaml:"statusId"`
}

