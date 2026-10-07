// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseAdminPasswordSource struct {
	// customer_managed_aws_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#customer_managed_aws_secret OdbAutonomousDatabase#customer_managed_aws_secret}
	CustomerManagedAwsSecret interface{} `field:"optional" json:"customerManagedAwsSecret" yaml:"customerManagedAwsSecret"`
}

