// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instance


type InstanceCapacityReservationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/instance#capacity_reservation_preference Instance#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/instance#capacity_reservation_target Instance#capacity_reservation_target}
	CapacityReservationTarget *InstanceCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

