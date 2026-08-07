// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigRagConfigKnowledgeBaseConfigRetrieveConfigKnowledgeBaseRetrievalConfiguration struct {
	// vector_search_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#vector_search_configuration BedrockEvaluationJob#vector_search_configuration}
	VectorSearchConfiguration interface{} `field:"optional" json:"vectorSearchConfiguration" yaml:"vectorSearchConfiguration"`
}

