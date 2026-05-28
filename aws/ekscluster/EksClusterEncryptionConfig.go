// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterEncryptionConfig struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/eks_cluster#provider EksCluster#provider}
	Provider *EksClusterEncryptionConfigProvider `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/eks_cluster#resources EksCluster#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

