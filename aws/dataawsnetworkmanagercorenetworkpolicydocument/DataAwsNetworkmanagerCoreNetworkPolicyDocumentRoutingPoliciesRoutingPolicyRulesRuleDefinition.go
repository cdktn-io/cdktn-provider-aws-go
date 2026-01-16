// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentRoutingPoliciesRoutingPolicyRulesRuleDefinition struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#action DataAwsNetworkmanagerCoreNetworkPolicyDocument#action}
	Action *DataAwsNetworkmanagerCoreNetworkPolicyDocumentRoutingPoliciesRoutingPolicyRulesRuleDefinitionAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#condition_logic DataAwsNetworkmanagerCoreNetworkPolicyDocument#condition_logic}.
	ConditionLogic *string `field:"optional" json:"conditionLogic" yaml:"conditionLogic"`
	// match_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#match_conditions DataAwsNetworkmanagerCoreNetworkPolicyDocument#match_conditions}
	MatchConditions interface{} `field:"optional" json:"matchConditions" yaml:"matchConditions"`
}

