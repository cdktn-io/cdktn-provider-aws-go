// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationMonitoringConfigurationS3LoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#enabled AthenaWorkgroup#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#kms_key AthenaWorkgroup#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#log_location AthenaWorkgroup#log_location}.
	LogLocation *string `field:"optional" json:"logLocation" yaml:"logLocation"`
}

