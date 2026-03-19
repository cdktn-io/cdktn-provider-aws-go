// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleFilter struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/s3_bucket_replication_configuration#and S3BucketReplicationConfigurationA#and}
	And *S3BucketReplicationConfigurationRuleFilterAnd `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/s3_bucket_replication_configuration#prefix S3BucketReplicationConfigurationA#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/s3_bucket_replication_configuration#tag S3BucketReplicationConfigurationA#tag}
	Tag *S3BucketReplicationConfigurationRuleFilterTag `field:"optional" json:"tag" yaml:"tag"`
}

