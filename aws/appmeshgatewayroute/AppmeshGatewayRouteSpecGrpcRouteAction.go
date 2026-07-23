// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecGrpcRouteAction struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/appmesh_gateway_route#target AppmeshGatewayRoute#target}
	Target *AppmeshGatewayRouteSpecGrpcRouteActionTarget `field:"required" json:"target" yaml:"target"`
}

