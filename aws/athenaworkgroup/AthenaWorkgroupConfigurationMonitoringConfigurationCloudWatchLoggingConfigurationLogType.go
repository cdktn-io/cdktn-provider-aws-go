// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationMonitoringConfigurationCloudWatchLoggingConfigurationLogType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#key AthenaWorkgroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#values AthenaWorkgroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

