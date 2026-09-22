// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClusterClientAuthentication struct {
	// mtls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/msk_replicator#mtls MskReplicator#mtls}
	Mtls *MskReplicatorKafkaClusterClientAuthenticationMtls `field:"optional" json:"mtls" yaml:"mtls"`
	// sasl_scram block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/msk_replicator#sasl_scram MskReplicator#sasl_scram}
	SaslScram *MskReplicatorKafkaClusterClientAuthenticationSaslScram `field:"optional" json:"saslScram" yaml:"saslScram"`
}

