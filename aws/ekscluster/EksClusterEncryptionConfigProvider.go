// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterEncryptionConfigProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/eks_cluster#key_arn EksCluster#key_arn}.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
}

