// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsec2hosts


type DataAwsEc2HostsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/data-sources/ec2_hosts#name DataAwsEc2Hosts#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/data-sources/ec2_hosts#values DataAwsEc2Hosts#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

