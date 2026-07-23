// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_type ObservabilityadminTelemetryRule#telemetry_type}.
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#allow_field_updates ObservabilityadminTelemetryRule#allow_field_updates}.
	AllowFieldUpdates interface{} `field:"optional" json:"allowFieldUpdates" yaml:"allowFieldUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#all_regions ObservabilityadminTelemetryRule#all_regions}.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#destination_configuration ObservabilityadminTelemetryRule#destination_configuration}
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#regions ObservabilityadminTelemetryRule#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#resource_type ObservabilityadminTelemetryRule#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#scope ObservabilityadminTelemetryRule#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#selection_criteria ObservabilityadminTelemetryRule#selection_criteria}.
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_source_types ObservabilityadminTelemetryRule#telemetry_source_types}.
	TelemetrySourceTypes *[]*string `field:"optional" json:"telemetrySourceTypes" yaml:"telemetrySourceTypes"`
}

