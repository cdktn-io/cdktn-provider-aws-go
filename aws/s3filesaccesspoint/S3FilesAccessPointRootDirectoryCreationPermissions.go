// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesaccesspoint


type S3FilesAccessPointRootDirectoryCreationPermissions struct {
	// Owner group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#owner_gid S3FilesAccessPoint#owner_gid}
	OwnerGid *float64 `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Owner user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#owner_uid S3FilesAccessPoint#owner_uid}
	OwnerUid *float64 `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// POSIX permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#permissions S3FilesAccessPoint#permissions}
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

