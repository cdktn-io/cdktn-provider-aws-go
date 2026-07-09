// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesaccesspoint


type S3FilesAccessPointPosixUser struct {
	// POSIX group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/s3files_access_point#gid S3FilesAccessPoint#gid}
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// POSIX user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/s3files_access_point#uid S3FilesAccessPoint#uid}
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/s3files_access_point#secondary_gids S3FilesAccessPoint#secondary_gids}
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

