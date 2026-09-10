// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClusterClientAuthenticationMtls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/msk_replicator#secret_arn MskReplicator#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

