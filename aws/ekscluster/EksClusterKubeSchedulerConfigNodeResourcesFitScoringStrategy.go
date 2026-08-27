// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeSchedulerConfigNodeResourcesFitScoringStrategy struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#resource EksCluster#resource}
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#type EksCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

