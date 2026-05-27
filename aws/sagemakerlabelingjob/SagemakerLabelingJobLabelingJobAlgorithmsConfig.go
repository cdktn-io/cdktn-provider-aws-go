// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobLabelingJobAlgorithmsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_labeling_job#labeling_job_algorithm_specification_arn SagemakerLabelingJob#labeling_job_algorithm_specification_arn}.
	LabelingJobAlgorithmSpecificationArn *string `field:"required" json:"labelingJobAlgorithmSpecificationArn" yaml:"labelingJobAlgorithmSpecificationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_labeling_job#initial_active_learning_model_arn SagemakerLabelingJob#initial_active_learning_model_arn}.
	InitialActiveLearningModelArn *string `field:"optional" json:"initialActiveLearningModelArn" yaml:"initialActiveLearningModelArn"`
	// labeling_job_resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_labeling_job#labeling_job_resource_config SagemakerLabelingJob#labeling_job_resource_config}
	LabelingJobResourceConfig interface{} `field:"optional" json:"labelingJobResourceConfig" yaml:"labelingJobResourceConfig"`
}

