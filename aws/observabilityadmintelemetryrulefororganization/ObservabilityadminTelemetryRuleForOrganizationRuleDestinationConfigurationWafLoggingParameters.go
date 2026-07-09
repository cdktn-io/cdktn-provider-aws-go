// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfigurationWafLoggingParameters struct {
	// logging_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#logging_filter ObservabilityadminTelemetryRuleForOrganization#logging_filter}
	LoggingFilter interface{} `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#log_type ObservabilityadminTelemetryRuleForOrganization#log_type}.
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#redacted_fields ObservabilityadminTelemetryRuleForOrganization#redacted_fields}
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

