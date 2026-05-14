// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsvolume


type FsxOpenzfsVolumeOriginSnapshot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/fsx_openzfs_volume#copy_strategy FsxOpenzfsVolume#copy_strategy}.
	CopyStrategy *string `field:"required" json:"copyStrategy" yaml:"copyStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/fsx_openzfs_volume#snapshot_arn FsxOpenzfsVolume#snapshot_arn}.
	SnapshotArn *string `field:"required" json:"snapshotArn" yaml:"snapshotArn"`
}

