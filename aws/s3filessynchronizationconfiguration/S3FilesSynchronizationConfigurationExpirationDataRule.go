// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filessynchronizationconfiguration


type S3FilesSynchronizationConfigurationExpirationDataRule struct {
	// Days after last access before data expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/s3files_synchronization_configuration#days_after_last_access S3FilesSynchronizationConfiguration#days_after_last_access}
	DaysAfterLastAccess *float64 `field:"required" json:"daysAfterLastAccess" yaml:"daysAfterLastAccess"`
}

