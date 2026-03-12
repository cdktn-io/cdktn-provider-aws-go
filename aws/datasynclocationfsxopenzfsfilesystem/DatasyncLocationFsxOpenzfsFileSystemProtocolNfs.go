// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxopenzfsfilesystem


type DatasyncLocationFsxOpenzfsFileSystemProtocolNfs struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/datasync_location_fsx_openzfs_file_system#mount_options DatasyncLocationFsxOpenzfsFileSystem#mount_options}
	MountOptions *DatasyncLocationFsxOpenzfsFileSystemProtocolNfsMountOptions `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

