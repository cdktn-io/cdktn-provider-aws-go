// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakedatalake


type SecuritylakeDataLakeConfigurationLifecycleConfigurationTransition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/securitylake_data_lake#days SecuritylakeDataLake#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/securitylake_data_lake#storage_class SecuritylakeDataLake#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

