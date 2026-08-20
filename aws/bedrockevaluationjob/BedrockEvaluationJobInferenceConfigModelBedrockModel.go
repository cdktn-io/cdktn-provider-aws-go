// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigModelBedrockModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/bedrock_evaluation_job#model_identifier BedrockEvaluationJob#model_identifier}.
	ModelIdentifier *string `field:"required" json:"modelIdentifier" yaml:"modelIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/bedrock_evaluation_job#inference_params BedrockEvaluationJob#inference_params}.
	InferenceParams *string `field:"optional" json:"inferenceParams" yaml:"inferenceParams"`
	// performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/bedrock_evaluation_job#performance_config BedrockEvaluationJob#performance_config}
	PerformanceConfig interface{} `field:"optional" json:"performanceConfig" yaml:"performanceConfig"`
}

