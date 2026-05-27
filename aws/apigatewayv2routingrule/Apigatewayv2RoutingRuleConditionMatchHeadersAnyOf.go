// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleConditionMatchHeadersAnyOf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/apigatewayv2_routing_rule#header Apigatewayv2RoutingRule#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/apigatewayv2_routing_rule#value_glob Apigatewayv2RoutingRule#value_glob}.
	ValueGlob *string `field:"required" json:"valueGlob" yaml:"valueGlob"`
}

