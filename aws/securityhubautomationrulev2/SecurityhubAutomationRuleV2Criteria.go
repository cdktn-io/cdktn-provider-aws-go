// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2Criteria struct {
	// JSON-encoded OCSF finding criteria for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/securityhub_automation_rule_v2#ocsf_finding_criteria_json SecurityhubAutomationRuleV2#ocsf_finding_criteria_json}
	OcsfFindingCriteriaJson *string `field:"required" json:"ocsfFindingCriteriaJson" yaml:"ocsfFindingCriteriaJson"`
}

