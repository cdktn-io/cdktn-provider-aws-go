// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsrdssnapshots


type DataAwsRdsSnapshotsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/data-sources/rds_snapshots#name DataAwsRdsSnapshots#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/data-sources/rds_snapshots#values DataAwsRdsSnapshots#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

