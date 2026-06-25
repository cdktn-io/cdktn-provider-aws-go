// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule#log_format ObservabilityadminTelemetryRule#log_format}.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule#max_aggregation_interval ObservabilityadminTelemetryRule#max_aggregation_interval}.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/observabilityadmin_telemetry_rule#traffic_type ObservabilityadminTelemetryRule#traffic_type}.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

