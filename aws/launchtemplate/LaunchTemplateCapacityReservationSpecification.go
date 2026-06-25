// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateCapacityReservationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/launch_template#capacity_reservation_preference LaunchTemplate#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/launch_template#capacity_reservation_target LaunchTemplate#capacity_reservation_target}
	CapacityReservationTarget *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

