// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2


type SecurityhubAutomationRuleV2Action struct {
	// The action type: FINDING_FIELDS_UPDATE or EXTERNAL_INTEGRATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/securityhub_automation_rule_v2#type SecurityhubAutomationRuleV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// external_integration_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/securityhub_automation_rule_v2#external_integration_configuration SecurityhubAutomationRuleV2#external_integration_configuration}
	ExternalIntegrationConfiguration interface{} `field:"optional" json:"externalIntegrationConfiguration" yaml:"externalIntegrationConfiguration"`
	// finding_fields_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/securityhub_automation_rule_v2#finding_fields_update SecurityhubAutomationRuleV2#finding_fields_update}
	FindingFieldsUpdate interface{} `field:"optional" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
}

