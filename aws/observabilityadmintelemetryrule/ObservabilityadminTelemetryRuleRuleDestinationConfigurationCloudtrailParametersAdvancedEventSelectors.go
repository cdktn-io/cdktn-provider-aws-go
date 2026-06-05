// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParametersAdvancedEventSelectors struct {
	// field_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/observabilityadmin_telemetry_rule#field_selectors ObservabilityadminTelemetryRule#field_selectors}
	FieldSelectors interface{} `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/observabilityadmin_telemetry_rule#name ObservabilityadminTelemetryRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

