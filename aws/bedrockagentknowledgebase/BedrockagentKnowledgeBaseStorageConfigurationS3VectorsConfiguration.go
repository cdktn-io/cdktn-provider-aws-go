// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseStorageConfigurationS3VectorsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#index_arn BedrockagentKnowledgeBase#index_arn}.
	IndexArn *string `field:"optional" json:"indexArn" yaml:"indexArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#index_name BedrockagentKnowledgeBase#index_name}.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#vector_bucket_arn BedrockagentKnowledgeBase#vector_bucket_arn}.
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}

