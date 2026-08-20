// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeSchedulerConfigNodeResourcesFit struct {
	// scoring_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/eks_cluster#scoring_strategy EksCluster#scoring_strategy}
	ScoringStrategy *EksClusterKubeSchedulerConfigNodeResourcesFitScoringStrategy `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
}

