// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_user_profile#default_ebs_volume_size_in_gb SagemakerUserProfile#default_ebs_volume_size_in_gb}.
	DefaultEbsVolumeSizeInGb *float64 `field:"required" json:"defaultEbsVolumeSizeInGb" yaml:"defaultEbsVolumeSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_user_profile#maximum_ebs_volume_size_in_gb SagemakerUserProfile#maximum_ebs_volume_size_in_gb}.
	MaximumEbsVolumeSizeInGb *float64 `field:"required" json:"maximumEbsVolumeSizeInGb" yaml:"maximumEbsVolumeSizeInGb"`
}

