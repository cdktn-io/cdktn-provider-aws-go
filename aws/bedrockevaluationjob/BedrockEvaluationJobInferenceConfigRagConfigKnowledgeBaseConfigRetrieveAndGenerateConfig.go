// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigRagConfigKnowledgeBaseConfigRetrieveAndGenerateConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/bedrock_evaluation_job#knowledge_base_id BedrockEvaluationJob#knowledge_base_id}.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/bedrock_evaluation_job#model_arn BedrockEvaluationJob#model_arn}.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// retrieval_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/bedrock_evaluation_job#retrieval_configuration BedrockEvaluationJob#retrieval_configuration}
	RetrievalConfiguration interface{} `field:"optional" json:"retrievalConfiguration" yaml:"retrievalConfiguration"`
}

