// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterOutpostConfigEtcdPlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/eks_cluster#spread_level EksCluster#spread_level}.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

