// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersLoggingFilterFiltersConditions struct {
	// action_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#action_condition ObservabilityadminTelemetryRuleForOrganization#action_condition}
	ActionCondition interface{} `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// label_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#label_name_condition ObservabilityadminTelemetryRuleForOrganization#label_name_condition}
	LabelNameCondition interface{} `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

