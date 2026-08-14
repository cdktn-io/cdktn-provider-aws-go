// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsfilesystem


type FsxOpenzfsFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/fsx_openzfs_file_system#iops FsxOpenzfsFileSystem#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/fsx_openzfs_file_system#mode FsxOpenzfsFileSystem#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

