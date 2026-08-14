// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeApiServerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/eks_cluster#event_ttl EksCluster#event_ttl}.
	EventTtl *string `field:"optional" json:"eventTtl" yaml:"eventTtl"`
	// service_node_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/eks_cluster#service_node_port_range EksCluster#service_node_port_range}
	ServiceNodePortRange *EksClusterKubeApiServerConfigServiceNodePortRange `field:"optional" json:"serviceNodePortRange" yaml:"serviceNodePortRange"`
}

