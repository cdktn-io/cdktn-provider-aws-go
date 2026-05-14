// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkacls


type DataAwsNetworkAclsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/network_acls#name DataAwsNetworkAcls#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/network_acls#values DataAwsNetworkAcls#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

