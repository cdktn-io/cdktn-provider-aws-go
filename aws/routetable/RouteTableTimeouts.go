// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package routetable


type RouteTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/route_table#create RouteTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/route_table#delete RouteTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/route_table#update RouteTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

