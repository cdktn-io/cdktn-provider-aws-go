// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMaintenanceConfigurationIcebergSnapshotManagementSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3tables_table#max_snapshot_age_hours S3TablesTable#max_snapshot_age_hours}.
	MaxSnapshotAgeHours *float64 `field:"optional" json:"maxSnapshotAgeHours" yaml:"maxSnapshotAgeHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/s3tables_table#min_snapshots_to_keep S3TablesTable#min_snapshots_to_keep}.
	MinSnapshotsToKeep *float64 `field:"optional" json:"minSnapshotsToKeep" yaml:"minSnapshotsToKeep"`
}

