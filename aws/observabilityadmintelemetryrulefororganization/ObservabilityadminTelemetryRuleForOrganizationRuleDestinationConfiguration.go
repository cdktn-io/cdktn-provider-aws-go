// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrulefororganization


type ObservabilityadminTelemetryRuleForOrganizationRuleDestinationConfiguration struct {
	// cloudtrail_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#cloudtrail_parameters ObservabilityadminTelemetryRuleForOrganization#cloudtrail_parameters}
	CloudtrailParameters interface{} `field:"optional" json:"cloudtrailParameters" yaml:"cloudtrailParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#destination_pattern ObservabilityadminTelemetryRuleForOrganization#destination_pattern}.
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#destination_type ObservabilityadminTelemetryRuleForOrganization#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// elb_load_balancer_logging_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#elb_load_balancer_logging_parameters ObservabilityadminTelemetryRuleForOrganization#elb_load_balancer_logging_parameters}
	ElbLoadBalancerLoggingParameters interface{} `field:"optional" json:"elbLoadBalancerLoggingParameters" yaml:"elbLoadBalancerLoggingParameters"`
	// log_delivery_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#log_delivery_parameters ObservabilityadminTelemetryRuleForOrganization#log_delivery_parameters}
	LogDeliveryParameters interface{} `field:"optional" json:"logDeliveryParameters" yaml:"logDeliveryParameters"`
	// msk_monitoring_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#msk_monitoring_parameters ObservabilityadminTelemetryRuleForOrganization#msk_monitoring_parameters}
	MskMonitoringParameters interface{} `field:"optional" json:"mskMonitoringParameters" yaml:"mskMonitoringParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#retention_in_days ObservabilityadminTelemetryRuleForOrganization#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// vpc_flow_log_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#vpc_flow_log_parameters ObservabilityadminTelemetryRuleForOrganization#vpc_flow_log_parameters}
	VpcFlowLogParameters interface{} `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
	// waf_logging_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#waf_logging_parameters ObservabilityadminTelemetryRuleForOrganization#waf_logging_parameters}
	WafLoggingParameters interface{} `field:"optional" json:"wafLoggingParameters" yaml:"wafLoggingParameters"`
}

