// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersLoggingFilterFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#behavior ObservabilityadminTelemetryRuleForOrganization#behavior}.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#conditions ObservabilityadminTelemetryRuleForOrganization#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#requirement ObservabilityadminTelemetryRuleForOrganization#requirement}.
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

