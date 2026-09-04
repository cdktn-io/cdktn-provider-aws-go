// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigRagConfigKnowledgeBaseConfig struct {
	// retrieve_and_generate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#retrieve_and_generate_config BedrockEvaluationJob#retrieve_and_generate_config}
	RetrieveAndGenerateConfig interface{} `field:"optional" json:"retrieveAndGenerateConfig" yaml:"retrieveAndGenerateConfig"`
	// retrieve_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#retrieve_config BedrockEvaluationJob#retrieve_config}
	RetrieveConfig interface{} `field:"optional" json:"retrieveConfig" yaml:"retrieveConfig"`
}

