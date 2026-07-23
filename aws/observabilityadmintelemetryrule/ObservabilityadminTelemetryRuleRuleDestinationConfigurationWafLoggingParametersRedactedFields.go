// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParametersRedactedFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#method ObservabilityadminTelemetryRule#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#query_string ObservabilityadminTelemetryRule#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#single_header ObservabilityadminTelemetryRule#single_header}
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/observabilityadmin_telemetry_rule#uri_path ObservabilityadminTelemetryRule#uri_path}.
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

