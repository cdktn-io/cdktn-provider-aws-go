// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationMonitoringConfiguration struct {
	// cloud_watch_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#cloud_watch_logging_configuration AthenaWorkgroup#cloud_watch_logging_configuration}
	CloudWatchLoggingConfiguration *AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration `field:"optional" json:"cloudWatchLoggingConfiguration" yaml:"cloudWatchLoggingConfiguration"`
	// managed_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#managed_logging_configuration AthenaWorkgroup#managed_logging_configuration}
	ManagedLoggingConfiguration *AthenaWorkgroupConfigurationMonitoringConfigurationManagedLoggingConfiguration `field:"optional" json:"managedLoggingConfiguration" yaml:"managedLoggingConfiguration"`
	// s3_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#s3_logging_configuration AthenaWorkgroup#s3_logging_configuration}
	S3LoggingConfiguration *AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration `field:"optional" json:"s3LoggingConfiguration" yaml:"s3LoggingConfiguration"`
}

