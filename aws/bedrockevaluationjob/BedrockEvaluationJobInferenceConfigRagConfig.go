// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigRagConfig struct {
	// knowledge_base_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#knowledge_base_config BedrockEvaluationJob#knowledge_base_config}
	KnowledgeBaseConfig interface{} `field:"optional" json:"knowledgeBaseConfig" yaml:"knowledgeBaseConfig"`
	// precomputed_rag_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#precomputed_rag_source_config BedrockEvaluationJob#precomputed_rag_source_config}
	PrecomputedRagSourceConfig interface{} `field:"optional" json:"precomputedRagSourceConfig" yaml:"precomputedRagSourceConfig"`
}

