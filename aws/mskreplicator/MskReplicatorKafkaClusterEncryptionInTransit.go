// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClusterEncryptionInTransit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/msk_replicator#root_ca_certificate MskReplicator#root_ca_certificate}.
	RootCaCertificate *string `field:"required" json:"rootCaCertificate" yaml:"rootCaCertificate"`
}

