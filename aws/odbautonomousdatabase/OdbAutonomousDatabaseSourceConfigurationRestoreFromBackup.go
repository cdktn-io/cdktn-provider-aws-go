// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseSourceConfigurationRestoreFromBackup struct {
	// ID of the Autonomous Database backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#autonomous_database_backup_id OdbAutonomousDatabase#autonomous_database_backup_id}
	AutonomousDatabaseBackupId *string `field:"required" json:"autonomousDatabaseBackupId" yaml:"autonomousDatabaseBackupId"`
	// Type of clone to create from the backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#clone_type OdbAutonomousDatabase#clone_type}
	CloneType *string `field:"required" json:"cloneType" yaml:"cloneType"`
	// Tablespace IDs to clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#clone_table_space_list OdbAutonomousDatabase#clone_table_space_list}
	CloneTableSpaceList *[]*float64 `field:"optional" json:"cloneTableSpaceList" yaml:"cloneTableSpaceList"`
}

