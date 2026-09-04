// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeApiServerConfigServiceNodePortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_cluster#max_port EksCluster#max_port}.
	MaxPort *float64 `field:"optional" json:"maxPort" yaml:"maxPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_cluster#min_port EksCluster#min_port}.
	MinPort *float64 `field:"optional" json:"minPort" yaml:"minPort"`
}

