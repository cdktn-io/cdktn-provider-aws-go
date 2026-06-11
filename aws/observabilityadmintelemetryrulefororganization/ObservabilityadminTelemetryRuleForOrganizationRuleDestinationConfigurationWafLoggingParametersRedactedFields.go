// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParametersRedactedFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#method ObservabilityadminTelemetryRuleForOrganization#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#query_string ObservabilityadminTelemetryRuleForOrganization#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#single_header ObservabilityadminTelemetryRuleForOrganization#single_header}
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#uri_path ObservabilityadminTelemetryRuleForOrganization#uri_path}.
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

