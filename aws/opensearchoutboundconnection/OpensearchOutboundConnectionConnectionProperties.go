// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchoutboundconnection


type OpensearchOutboundConnectionConnectionProperties struct {
	// cross_cluster_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/opensearch_outbound_connection#cross_cluster_search OpensearchOutboundConnection#cross_cluster_search}
	CrossClusterSearch *OpensearchOutboundConnectionConnectionPropertiesCrossClusterSearch `field:"optional" json:"crossClusterSearch" yaml:"crossClusterSearch"`
}

