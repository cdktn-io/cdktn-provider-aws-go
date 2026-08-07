// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_knowledge_base#embedding_model_arn BedrockagentKnowledgeBase#embedding_model_arn}.
	EmbeddingModelArn *string `field:"optional" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// embedding_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_knowledge_base#embedding_model_configuration BedrockagentKnowledgeBase#embedding_model_configuration}
	EmbeddingModelConfiguration interface{} `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_knowledge_base#embedding_model_type BedrockagentKnowledgeBase#embedding_model_type}.
	EmbeddingModelType *string `field:"optional" json:"embeddingModelType" yaml:"embeddingModelType"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_knowledge_base#server_side_encryption_configuration BedrockagentKnowledgeBase#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

