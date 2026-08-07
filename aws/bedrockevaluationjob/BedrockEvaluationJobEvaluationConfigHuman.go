// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigHuman struct {
	// custom_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#custom_metric BedrockEvaluationJob#custom_metric}
	CustomMetric interface{} `field:"optional" json:"customMetric" yaml:"customMetric"`
	// dataset_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#dataset_metric_config BedrockEvaluationJob#dataset_metric_config}
	DatasetMetricConfig interface{} `field:"optional" json:"datasetMetricConfig" yaml:"datasetMetricConfig"`
	// human_workflow_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#human_workflow_config BedrockEvaluationJob#human_workflow_config}
	HumanWorkflowConfig interface{} `field:"optional" json:"humanWorkflowConfig" yaml:"humanWorkflowConfig"`
}

