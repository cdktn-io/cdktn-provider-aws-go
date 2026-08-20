// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2policy


type Resiliencehubv2PolicyDataRecovery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_policy#time_between_backups_in_minutes Resiliencehubv2Policy#time_between_backups_in_minutes}.
	TimeBetweenBackupsInMinutes *float64 `field:"required" json:"timeBetweenBackupsInMinutes" yaml:"timeBetweenBackupsInMinutes"`
}

