// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filessynchronizationconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3FilesSynchronizationConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// File system ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3files_synchronization_configuration#file_system_id S3FilesSynchronizationConfiguration#file_system_id}
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// expiration_data_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3files_synchronization_configuration#expiration_data_rule S3FilesSynchronizationConfiguration#expiration_data_rule}
	ExpirationDataRule interface{} `field:"optional" json:"expirationDataRule" yaml:"expirationDataRule"`
	// import_data_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3files_synchronization_configuration#import_data_rule S3FilesSynchronizationConfiguration#import_data_rule}
	ImportDataRule interface{} `field:"optional" json:"importDataRule" yaml:"importDataRule"`
	// Latest version number for optimistic locking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3files_synchronization_configuration#latest_version_number S3FilesSynchronizationConfiguration#latest_version_number}
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3files_synchronization_configuration#region S3FilesSynchronizationConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

