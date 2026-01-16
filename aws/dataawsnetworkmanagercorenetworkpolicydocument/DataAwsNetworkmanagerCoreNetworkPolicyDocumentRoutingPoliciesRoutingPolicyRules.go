// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentRoutingPoliciesRoutingPolicyRules struct {
	// rule_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#rule_definition DataAwsNetworkmanagerCoreNetworkPolicyDocument#rule_definition}
	RuleDefinition *DataAwsNetworkmanagerCoreNetworkPolicyDocumentRoutingPoliciesRoutingPolicyRulesRuleDefinition `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#rule_number DataAwsNetworkmanagerCoreNetworkPolicyDocument#rule_number}.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
}

