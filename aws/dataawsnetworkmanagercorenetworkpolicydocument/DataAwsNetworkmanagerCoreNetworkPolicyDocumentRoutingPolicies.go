// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentRoutingPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_direction DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_direction}.
	RoutingPolicyDirection *string `field:"required" json:"routingPolicyDirection" yaml:"routingPolicyDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_name DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_name}.
	RoutingPolicyName *string `field:"required" json:"routingPolicyName" yaml:"routingPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_number DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_number}.
	RoutingPolicyNumber *float64 `field:"required" json:"routingPolicyNumber" yaml:"routingPolicyNumber"`
	// routing_policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_rules DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_rules}
	RoutingPolicyRules interface{} `field:"required" json:"routingPolicyRules" yaml:"routingPolicyRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_description DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_description}.
	RoutingPolicyDescription *string `field:"optional" json:"routingPolicyDescription" yaml:"routingPolicyDescription"`
}

