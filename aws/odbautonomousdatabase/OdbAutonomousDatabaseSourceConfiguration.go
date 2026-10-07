// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseSourceConfiguration struct {
	// clone_to_refreshable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#clone_to_refreshable OdbAutonomousDatabase#clone_to_refreshable}
	CloneToRefreshable interface{} `field:"optional" json:"cloneToRefreshable" yaml:"cloneToRefreshable"`
	// cross_region_data_guard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#cross_region_data_guard OdbAutonomousDatabase#cross_region_data_guard}
	CrossRegionDataGuard interface{} `field:"optional" json:"crossRegionDataGuard" yaml:"crossRegionDataGuard"`
	// cross_region_disaster_recovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#cross_region_disaster_recovery OdbAutonomousDatabase#cross_region_disaster_recovery}
	CrossRegionDisasterRecovery interface{} `field:"optional" json:"crossRegionDisasterRecovery" yaml:"crossRegionDisasterRecovery"`
	// database_clone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#database_clone OdbAutonomousDatabase#database_clone}
	DatabaseClone interface{} `field:"optional" json:"databaseClone" yaml:"databaseClone"`
	// point_in_time_restore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#point_in_time_restore OdbAutonomousDatabase#point_in_time_restore}
	PointInTimeRestore interface{} `field:"optional" json:"pointInTimeRestore" yaml:"pointInTimeRestore"`
	// restore_from_backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#restore_from_backup OdbAutonomousDatabase#restore_from_backup}
	RestoreFromBackup interface{} `field:"optional" json:"restoreFromBackup" yaml:"restoreFromBackup"`
}

