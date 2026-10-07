// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuard struct {
	// ARN of the source Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#source_autonomous_database_arn OdbAutonomousDatabase#source_autonomous_database_arn}
	SourceAutonomousDatabaseArn *string `field:"required" json:"sourceAutonomousDatabaseArn" yaml:"sourceAutonomousDatabaseArn"`
}

