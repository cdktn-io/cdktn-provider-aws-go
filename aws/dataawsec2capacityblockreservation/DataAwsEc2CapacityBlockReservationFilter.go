// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsec2capacityblockreservation


type DataAwsEc2CapacityBlockReservationFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_capacity_block_reservation#name DataAwsEc2CapacityBlockReservation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_capacity_block_reservation#values DataAwsEc2CapacityBlockReservation#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

