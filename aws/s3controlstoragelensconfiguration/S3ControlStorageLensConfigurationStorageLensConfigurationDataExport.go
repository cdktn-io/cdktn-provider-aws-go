// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationDataExport struct {
	// cloud_watch_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/s3control_storage_lens_configuration#cloud_watch_metrics S3ControlStorageLensConfiguration#cloud_watch_metrics}
	CloudWatchMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetrics `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/s3control_storage_lens_configuration#s3_bucket_destination S3ControlStorageLensConfiguration#s3_bucket_destination}
	S3BucketDestination *S3ControlStorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestination `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// storage_lens_table_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/s3control_storage_lens_configuration#storage_lens_table_destination S3ControlStorageLensConfiguration#storage_lens_table_destination}
	StorageLensTableDestination *S3ControlStorageLensConfigurationStorageLensConfigurationDataExportStorageLensTableDestination `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

