// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigAutomated struct {
	// custom_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#custom_metric_config BedrockEvaluationJob#custom_metric_config}
	CustomMetricConfig interface{} `field:"optional" json:"customMetricConfig" yaml:"customMetricConfig"`
	// dataset_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#dataset_metric_config BedrockEvaluationJob#dataset_metric_config}
	DatasetMetricConfig interface{} `field:"optional" json:"datasetMetricConfig" yaml:"datasetMetricConfig"`
	// evaluator_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#evaluator_model_config BedrockEvaluationJob#evaluator_model_config}
	EvaluatorModelConfig interface{} `field:"optional" json:"evaluatorModelConfig" yaml:"evaluatorModelConfig"`
}

