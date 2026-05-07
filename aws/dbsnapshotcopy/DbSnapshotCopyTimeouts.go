// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dbsnapshotcopy


type DbSnapshotCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/db_snapshot_copy#create DbSnapshotCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

