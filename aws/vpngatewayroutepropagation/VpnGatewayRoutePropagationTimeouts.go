// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpngatewayroutepropagation


type VpnGatewayRoutePropagationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/vpn_gateway_route_propagation#create VpnGatewayRoutePropagation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/vpn_gateway_route_propagation#delete VpnGatewayRoutePropagation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

