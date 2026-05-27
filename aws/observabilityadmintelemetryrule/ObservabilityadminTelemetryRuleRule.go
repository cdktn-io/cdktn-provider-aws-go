// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_type ObservabilityadminTelemetryRule#telemetry_type}.
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/observabilityadmin_telemetry_rule#resource_type ObservabilityadminTelemetryRule#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

