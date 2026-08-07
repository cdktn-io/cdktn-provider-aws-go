// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigAutomatedDatasetMetricConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#metric_names BedrockEvaluationJob#metric_names}.
	MetricNames *[]*string `field:"required" json:"metricNames" yaml:"metricNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#task_type BedrockEvaluationJob#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrock_evaluation_job#dataset BedrockEvaluationJob#dataset}
	Dataset interface{} `field:"optional" json:"dataset" yaml:"dataset"`
}

