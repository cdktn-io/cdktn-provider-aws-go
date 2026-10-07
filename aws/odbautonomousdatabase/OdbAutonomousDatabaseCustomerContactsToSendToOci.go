// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseCustomerContactsToSendToOci struct {
	// Email address of the customer contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#email OdbAutonomousDatabase#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

