// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsec2servicelinkvirtualinterfaces


type DataAwsEc2ServiceLinkVirtualInterfacesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/data-sources/ec2_service_link_virtual_interfaces#name DataAwsEc2ServiceLinkVirtualInterfaces#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/data-sources/ec2_service_link_virtual_interfaces#values DataAwsEc2ServiceLinkVirtualInterfaces#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

