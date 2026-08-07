// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3objectcopy


type S3ObjectCopyOverrideProvider struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/s3_object_copy#default_tags S3ObjectCopy#default_tags}
	DefaultTags *S3ObjectCopyOverrideProviderDefaultTags `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

