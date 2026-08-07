// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsec2servicelinkvirtualinterface


type DataAwsEc2ServiceLinkVirtualInterfaceFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/data-sources/ec2_service_link_virtual_interface#name DataAwsEc2ServiceLinkVirtualInterface#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/data-sources/ec2_service_link_virtual_interface#values DataAwsEc2ServiceLinkVirtualInterface#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

