// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseSourceConfigurationPointInTimeRestore struct {
	// Type of clone to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#clone_type OdbAutonomousDatabase#clone_type}
	CloneType *string `field:"required" json:"cloneType" yaml:"cloneType"`
	// ID of the source Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#source_autonomous_database_id OdbAutonomousDatabase#source_autonomous_database_id}
	SourceAutonomousDatabaseId *string `field:"required" json:"sourceAutonomousDatabaseId" yaml:"sourceAutonomousDatabaseId"`
	// Tablespace IDs to clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#clone_table_space_list OdbAutonomousDatabase#clone_table_space_list}
	CloneTableSpaceList *[]*float64 `field:"optional" json:"cloneTableSpaceList" yaml:"cloneTableSpaceList"`
	// Date and time to which the database is restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#timestamp OdbAutonomousDatabase#timestamp}
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
	// Whether to use the latest available backup timestamp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#use_latest_available_backup_timestamp OdbAutonomousDatabase#use_latest_available_backup_timestamp}
	UseLatestAvailableBackupTimestamp interface{} `field:"optional" json:"useLatestAvailableBackupTimestamp" yaml:"useLatestAvailableBackupTimestamp"`
}

