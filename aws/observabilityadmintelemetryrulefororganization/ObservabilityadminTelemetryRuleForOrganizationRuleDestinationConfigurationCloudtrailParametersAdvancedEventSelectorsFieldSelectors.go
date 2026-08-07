// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationCloudtrailParametersAdvancedEventSelectorsFieldSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#field ObservabilityadminTelemetryRuleForOrganization#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#ends_with ObservabilityadminTelemetryRuleForOrganization#ends_with}.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#equals ObservabilityadminTelemetryRuleForOrganization#equals}.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#not_ends_with ObservabilityadminTelemetryRuleForOrganization#not_ends_with}.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#not_equals ObservabilityadminTelemetryRuleForOrganization#not_equals}.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#not_starts_with ObservabilityadminTelemetryRuleForOrganization#not_starts_with}.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#starts_with ObservabilityadminTelemetryRuleForOrganization#starts_with}.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

