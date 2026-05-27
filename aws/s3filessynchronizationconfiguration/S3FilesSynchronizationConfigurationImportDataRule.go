// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filessynchronizationconfiguration


type S3FilesSynchronizationConfigurationImportDataRule struct {
	// S3 prefix for import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/s3files_synchronization_configuration#prefix S3FilesSynchronizationConfiguration#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Maximum file size to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/s3files_synchronization_configuration#size_less_than S3FilesSynchronizationConfiguration#size_less_than}
	SizeLessThan *float64 `field:"required" json:"sizeLessThan" yaml:"sizeLessThan"`
	// Import trigger type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/s3files_synchronization_configuration#trigger S3FilesSynchronizationConfiguration#trigger}
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
}

