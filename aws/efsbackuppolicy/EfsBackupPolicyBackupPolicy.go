// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsbackuppolicy


type EfsBackupPolicyBackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/efs_backup_policy#status EfsBackupPolicy#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

