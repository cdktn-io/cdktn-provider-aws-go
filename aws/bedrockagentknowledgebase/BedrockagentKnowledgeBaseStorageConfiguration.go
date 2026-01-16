// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseStorageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// mongo_db_atlas_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#mongo_db_atlas_configuration BedrockagentKnowledgeBase#mongo_db_atlas_configuration}
	MongoDbAtlasConfiguration interface{} `field:"optional" json:"mongoDbAtlasConfiguration" yaml:"mongoDbAtlasConfiguration"`
	// neptune_analytics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#neptune_analytics_configuration BedrockagentKnowledgeBase#neptune_analytics_configuration}
	NeptuneAnalyticsConfiguration interface{} `field:"optional" json:"neptuneAnalyticsConfiguration" yaml:"neptuneAnalyticsConfiguration"`
	// opensearch_managed_cluster_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#opensearch_managed_cluster_configuration BedrockagentKnowledgeBase#opensearch_managed_cluster_configuration}
	OpensearchManagedClusterConfiguration interface{} `field:"optional" json:"opensearchManagedClusterConfiguration" yaml:"opensearchManagedClusterConfiguration"`
	// opensearch_serverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#opensearch_serverless_configuration BedrockagentKnowledgeBase#opensearch_serverless_configuration}
	OpensearchServerlessConfiguration interface{} `field:"optional" json:"opensearchServerlessConfiguration" yaml:"opensearchServerlessConfiguration"`
	// pinecone_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#pinecone_configuration BedrockagentKnowledgeBase#pinecone_configuration}
	PineconeConfiguration interface{} `field:"optional" json:"pineconeConfiguration" yaml:"pineconeConfiguration"`
	// rds_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#rds_configuration BedrockagentKnowledgeBase#rds_configuration}
	RdsConfiguration interface{} `field:"optional" json:"rdsConfiguration" yaml:"rdsConfiguration"`
	// redis_enterprise_cloud_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#redis_enterprise_cloud_configuration BedrockagentKnowledgeBase#redis_enterprise_cloud_configuration}
	RedisEnterpriseCloudConfiguration interface{} `field:"optional" json:"redisEnterpriseCloudConfiguration" yaml:"redisEnterpriseCloudConfiguration"`
	// s3_vectors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#s3_vectors_configuration BedrockagentKnowledgeBase#s3_vectors_configuration}
	S3VectorsConfiguration interface{} `field:"optional" json:"s3VectorsConfiguration" yaml:"s3VectorsConfiguration"`
}

