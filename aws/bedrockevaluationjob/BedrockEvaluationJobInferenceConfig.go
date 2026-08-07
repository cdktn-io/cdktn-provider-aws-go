// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfig struct {
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#model BedrockEvaluationJob#model}
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// rag_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#rag_config BedrockEvaluationJob#rag_config}
	RagConfig interface{} `field:"optional" json:"ragConfig" yaml:"ragConfig"`
}

