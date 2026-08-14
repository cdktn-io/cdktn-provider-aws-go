// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDataset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#name BedrockEvaluationJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// dataset_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrock_evaluation_job#dataset_location BedrockEvaluationJob#dataset_location}
	DatasetLocation interface{} `field:"optional" json:"datasetLocation" yaml:"datasetLocation"`
}

