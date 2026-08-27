// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigRagConfigKnowledgeBaseConfigRetrieveConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_id BedrockEvaluationJob#knowledge_base_id}.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// knowledge_base_retrieval_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_retrieval_configuration BedrockEvaluationJob#knowledge_base_retrieval_configuration}
	KnowledgeBaseRetrievalConfiguration interface{} `field:"optional" json:"knowledgeBaseRetrievalConfiguration" yaml:"knowledgeBaseRetrievalConfiguration"`
}

