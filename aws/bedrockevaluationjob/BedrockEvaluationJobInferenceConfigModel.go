// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobInferenceConfigModel struct {
	// bedrock_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#bedrock_model BedrockEvaluationJob#bedrock_model}
	BedrockModel interface{} `field:"optional" json:"bedrockModel" yaml:"bedrockModel"`
	// precomputed_inference_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#precomputed_inference_source BedrockEvaluationJob#precomputed_inference_source}
	PrecomputedInferenceSource interface{} `field:"optional" json:"precomputedInferenceSource" yaml:"precomputedInferenceSource"`
}

