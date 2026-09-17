// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClusterClientAuthenticationSaslScram struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#mechanism MskReplicator#mechanism}.
	Mechanism *string `field:"required" json:"mechanism" yaml:"mechanism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#secret_arn MskReplicator#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

