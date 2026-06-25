// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#telemetry_type ObservabilityadminTelemetryRuleForOrganization#telemetry_type}.
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#allow_field_updates ObservabilityadminTelemetryRuleForOrganization#allow_field_updates}.
	AllowFieldUpdates interface{} `field:"optional" json:"allowFieldUpdates" yaml:"allowFieldUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#all_regions ObservabilityadminTelemetryRuleForOrganization#all_regions}.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#destination_configuration ObservabilityadminTelemetryRuleForOrganization#destination_configuration}
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#regions ObservabilityadminTelemetryRuleForOrganization#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#resource_type ObservabilityadminTelemetryRuleForOrganization#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#scope ObservabilityadminTelemetryRuleForOrganization#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#selection_criteria ObservabilityadminTelemetryRuleForOrganization#selection_criteria}.
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#telemetry_source_types ObservabilityadminTelemetryRuleForOrganization#telemetry_source_types}.
	TelemetrySourceTypes *[]*string `field:"optional" json:"telemetrySourceTypes" yaml:"telemetrySourceTypes"`
}

