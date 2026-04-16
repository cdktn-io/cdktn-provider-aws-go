// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationQueryResultsS3AccessGrantsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/athena_workgroup#authentication_type AthenaWorkgroup#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/athena_workgroup#enable_s3_access_grants AthenaWorkgroup#enable_s3_access_grants}.
	EnableS3AccessGrants interface{} `field:"required" json:"enableS3AccessGrants" yaml:"enableS3AccessGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/athena_workgroup#create_user_level_prefix AthenaWorkgroup#create_user_level_prefix}.
	CreateUserLevelPrefix interface{} `field:"optional" json:"createUserLevelPrefix" yaml:"createUserLevelPrefix"`
}

