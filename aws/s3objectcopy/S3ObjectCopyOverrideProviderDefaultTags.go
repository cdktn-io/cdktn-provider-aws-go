// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3objectcopy


type S3ObjectCopyOverrideProviderDefaultTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/s3_object_copy#tags S3ObjectCopy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

