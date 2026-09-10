// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClusterApacheKafkaCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/msk_replicator#apache_kafka_cluster_id MskReplicator#apache_kafka_cluster_id}.
	ApacheKafkaClusterId *string `field:"required" json:"apacheKafkaClusterId" yaml:"apacheKafkaClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/msk_replicator#bootstrap_broker_string MskReplicator#bootstrap_broker_string}.
	BootstrapBrokerString *string `field:"required" json:"bootstrapBrokerString" yaml:"bootstrapBrokerString"`
}

