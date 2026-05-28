// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryption struct {
	// sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/s3control_storage_lens_configuration#sse_kms S3ControlStorageLensConfiguration#sse_kms}
	SseKms *S3ControlStorageLensConfigurationStorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKms `field:"optional" json:"sseKms" yaml:"sseKms"`
	// sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/s3control_storage_lens_configuration#sse_s3 S3ControlStorageLensConfiguration#sse_s3}
	SseS3 interface{} `field:"optional" json:"sseS3" yaml:"sseS3"`
}

