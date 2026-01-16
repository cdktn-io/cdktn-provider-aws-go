// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationCustomerContentEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/athena_workgroup#kms_key AthenaWorkgroup#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

