// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigRagConfigPrecomputedRagSourceConfig struct {
	// retrieve_and_generate_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#retrieve_and_generate_source_config BedrockEvaluationJob#retrieve_and_generate_source_config}
	RetrieveAndGenerateSourceConfig interface{} `field:"optional" json:"retrieveAndGenerateSourceConfig" yaml:"retrieveAndGenerateSourceConfig"`
	// retrieve_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#retrieve_source_config BedrockEvaluationJob#retrieve_source_config}
	RetrieveSourceConfig interface{} `field:"optional" json:"retrieveSourceConfig" yaml:"retrieveSourceConfig"`
}

