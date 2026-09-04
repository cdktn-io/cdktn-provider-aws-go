// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationVpcFlowLogParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#log_format ObservabilityadminTelemetryRuleForOrganization#log_format}.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#max_aggregation_interval ObservabilityadminTelemetryRuleForOrganization#max_aggregation_interval}.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#traffic_type ObservabilityadminTelemetryRuleForOrganization#traffic_type}.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

