// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaCluster struct {
	// amazon_msk_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#amazon_msk_cluster MskReplicator#amazon_msk_cluster}
	AmazonMskCluster *MskReplicatorKafkaClusterAmazonMskCluster `field:"optional" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// apache_kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#apache_kafka_cluster MskReplicator#apache_kafka_cluster}
	ApacheKafkaCluster *MskReplicatorKafkaClusterApacheKafkaCluster `field:"optional" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
	// client_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#client_authentication MskReplicator#client_authentication}
	ClientAuthentication *MskReplicatorKafkaClusterClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// encryption_in_transit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#encryption_in_transit MskReplicator#encryption_in_transit}
	EncryptionInTransit *MskReplicatorKafkaClusterEncryptionInTransit `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/msk_replicator#vpc_config MskReplicator#vpc_config}
	VpcConfig *MskReplicatorKafkaClusterVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

