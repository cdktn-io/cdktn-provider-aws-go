// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorReplicationInfoListStruct struct {
	// consumer_group_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#consumer_group_replication MskReplicator#consumer_group_replication}
	ConsumerGroupReplication interface{} `field:"required" json:"consumerGroupReplication" yaml:"consumerGroupReplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#target_compression_type MskReplicator#target_compression_type}.
	TargetCompressionType *string `field:"required" json:"targetCompressionType" yaml:"targetCompressionType"`
	// topic_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#topic_replication MskReplicator#topic_replication}
	TopicReplication interface{} `field:"required" json:"topicReplication" yaml:"topicReplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#source_kafka_cluster_arn MskReplicator#source_kafka_cluster_arn}.
	SourceKafkaClusterArn *string `field:"optional" json:"sourceKafkaClusterArn" yaml:"sourceKafkaClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#source_kafka_cluster_id MskReplicator#source_kafka_cluster_id}.
	SourceKafkaClusterId *string `field:"optional" json:"sourceKafkaClusterId" yaml:"sourceKafkaClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#target_kafka_cluster_arn MskReplicator#target_kafka_cluster_arn}.
	TargetKafkaClusterArn *string `field:"optional" json:"targetKafkaClusterArn" yaml:"targetKafkaClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#target_kafka_cluster_id MskReplicator#target_kafka_cluster_id}.
	TargetKafkaClusterId *string `field:"optional" json:"targetKafkaClusterId" yaml:"targetKafkaClusterId"`
}

