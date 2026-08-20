// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorLogDelivery struct {
	// replicator_log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/msk_replicator#replicator_log_delivery MskReplicator#replicator_log_delivery}
	ReplicatorLogDelivery *MskReplicatorLogDeliveryReplicatorLogDelivery `field:"optional" json:"replicatorLogDelivery" yaml:"replicatorLogDelivery"`
}

