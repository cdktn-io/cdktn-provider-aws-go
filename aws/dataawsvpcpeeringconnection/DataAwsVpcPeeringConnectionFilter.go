// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcpeeringconnection


type DataAwsVpcPeeringConnectionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/data-sources/vpc_peering_connection#name DataAwsVpcPeeringConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/data-sources/vpc_peering_connection#values DataAwsVpcPeeringConnection#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

