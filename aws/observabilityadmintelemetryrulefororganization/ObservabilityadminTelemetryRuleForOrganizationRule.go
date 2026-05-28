// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#telemetry_type ObservabilityadminTelemetryRuleForOrganization#telemetry_type}.
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#resource_type ObservabilityadminTelemetryRuleForOrganization#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

