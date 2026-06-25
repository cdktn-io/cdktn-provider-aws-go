// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleCondition struct {
	// match_base_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/apigatewayv2_routing_rule#match_base_paths Apigatewayv2RoutingRule#match_base_paths}
	MatchBasePaths interface{} `field:"optional" json:"matchBasePaths" yaml:"matchBasePaths"`
	// match_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/apigatewayv2_routing_rule#match_headers Apigatewayv2RoutingRule#match_headers}
	MatchHeaders interface{} `field:"optional" json:"matchHeaders" yaml:"matchHeaders"`
}

