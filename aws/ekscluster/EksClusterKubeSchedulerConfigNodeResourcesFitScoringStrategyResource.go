// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeSchedulerConfigNodeResourcesFitScoringStrategyResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/eks_cluster#name EksCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/eks_cluster#weight EksCluster#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

