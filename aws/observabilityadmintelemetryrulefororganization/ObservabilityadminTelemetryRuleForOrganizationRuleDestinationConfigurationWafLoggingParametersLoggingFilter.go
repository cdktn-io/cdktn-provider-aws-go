// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersLoggingFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#default_behavior ObservabilityadminTelemetryRuleForOrganization#default_behavior}.
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#filters ObservabilityadminTelemetryRuleForOrganization#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

