// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationAccessControlTranslation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/s3_bucket_replication_configuration#owner S3BucketReplicationConfigurationA#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

