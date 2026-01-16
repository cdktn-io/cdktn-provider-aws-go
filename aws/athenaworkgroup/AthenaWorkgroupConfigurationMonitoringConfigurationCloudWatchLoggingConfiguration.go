// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#enabled AthenaWorkgroup#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#log_group AthenaWorkgroup#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#log_stream_name_prefix AthenaWorkgroup#log_stream_name_prefix}.
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// log_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#log_type AthenaWorkgroup#log_type}
	LogType interface{} `field:"optional" json:"logType" yaml:"logType"`
}

