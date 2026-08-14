// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigAutomatedCustomMetricConfig struct {
	// custom_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#custom_metric BedrockEvaluationJob#custom_metric}
	CustomMetric interface{} `field:"optional" json:"customMetric" yaml:"customMetric"`
	// evaluator_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#evaluator_model_config BedrockEvaluationJob#evaluator_model_config}
	EvaluatorModelConfig interface{} `field:"optional" json:"evaluatorModelConfig" yaml:"evaluatorModelConfig"`
}

