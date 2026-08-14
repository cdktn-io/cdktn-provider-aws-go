// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewayconnectpeer


type DataAwsEc2TransitGatewayConnectPeerFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/data-sources/ec2_transit_gateway_connect_peer#name DataAwsEc2TransitGatewayConnectPeer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/data-sources/ec2_transit_gateway_connect_peer#values DataAwsEc2TransitGatewayConnectPeer#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

