// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsEdgeLocationAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#edge_location DataAwsNetworkmanagerCoreNetworkPolicyDocument#edge_location}.
	EdgeLocation *string `field:"required" json:"edgeLocation" yaml:"edgeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#peer_edge_location DataAwsNetworkmanagerCoreNetworkPolicyDocument#peer_edge_location}.
	PeerEdgeLocation *string `field:"required" json:"peerEdgeLocation" yaml:"peerEdgeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_names DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_names}.
	RoutingPolicyNames *[]*string `field:"required" json:"routingPolicyNames" yaml:"routingPolicyNames"`
}

