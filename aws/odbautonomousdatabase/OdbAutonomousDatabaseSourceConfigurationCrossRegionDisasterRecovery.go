// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecovery struct {
	// Type of remote disaster recovery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#remote_disaster_recovery_type OdbAutonomousDatabase#remote_disaster_recovery_type}
	RemoteDisasterRecoveryType *string `field:"required" json:"remoteDisasterRecoveryType" yaml:"remoteDisasterRecoveryType"`
	// ARN of the source Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#source_autonomous_database_arn OdbAutonomousDatabase#source_autonomous_database_arn}
	SourceAutonomousDatabaseArn *string `field:"required" json:"sourceAutonomousDatabaseArn" yaml:"sourceAutonomousDatabaseArn"`
	// Whether automatic backups are replicated to the disaster recovery database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_replicate_automatic_backups OdbAutonomousDatabase#is_replicate_automatic_backups}
	IsReplicateAutomaticBackups interface{} `field:"optional" json:"isReplicateAutomaticBackups" yaml:"isReplicateAutomaticBackups"`
}

