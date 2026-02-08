// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsnatgateways


type DataAwsNatGatewaysFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/data-sources/nat_gateways#name DataAwsNatGateways#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/data-sources/nat_gateways#values DataAwsNatGateways#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

