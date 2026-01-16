// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#action DataAwsNetworkmanagerCoreNetworkPolicyDocument#action}
	Action *DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentRoutingPolicyRulesAction `field:"required" json:"action" yaml:"action"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#conditions DataAwsNetworkmanagerCoreNetworkPolicyDocument#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#rule_number DataAwsNetworkmanagerCoreNetworkPolicyDocument#rule_number}.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#description DataAwsNetworkmanagerCoreNetworkPolicyDocument#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#edge_locations DataAwsNetworkmanagerCoreNetworkPolicyDocument#edge_locations}.
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
}

