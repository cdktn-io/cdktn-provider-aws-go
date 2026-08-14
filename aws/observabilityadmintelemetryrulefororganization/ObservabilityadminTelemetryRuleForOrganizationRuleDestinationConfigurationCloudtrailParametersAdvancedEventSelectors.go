// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationCloudtrailParametersAdvancedEventSelectors struct {
	// field_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#field_selectors ObservabilityadminTelemetryRuleForOrganization#field_selectors}
	FieldSelectors interface{} `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#name ObservabilityadminTelemetryRuleForOrganization#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

