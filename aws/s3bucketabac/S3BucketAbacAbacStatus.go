// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucketabac


type S3BucketAbacAbacStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/s3_bucket_abac#status S3BucketAbac#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

