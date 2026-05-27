// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbdbinstance


type TimestreaminfluxdbDbInstanceMaintenanceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/timestreaminfluxdb_db_instance#preferred_maintenance_window TimestreaminfluxdbDbInstance#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"required" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/timestreaminfluxdb_db_instance#timezone TimestreaminfluxdbDbInstance#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

