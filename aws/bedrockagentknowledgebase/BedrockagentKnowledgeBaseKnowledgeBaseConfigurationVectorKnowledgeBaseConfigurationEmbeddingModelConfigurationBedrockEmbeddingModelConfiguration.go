// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagent_knowledge_base#dimensions BedrockagentKnowledgeBase#dimensions}.
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagent_knowledge_base#embedding_data_type BedrockagentKnowledgeBase#embedding_data_type}.
	EmbeddingDataType *string `field:"optional" json:"embeddingDataType" yaml:"embeddingDataType"`
}

