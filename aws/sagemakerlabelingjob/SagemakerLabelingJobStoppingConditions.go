// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobStoppingConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/sagemaker_labeling_job#max_human_labeled_object_count SagemakerLabelingJob#max_human_labeled_object_count}.
	MaxHumanLabeledObjectCount *float64 `field:"optional" json:"maxHumanLabeledObjectCount" yaml:"maxHumanLabeledObjectCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/sagemaker_labeling_job#max_percentage_of_input_dataset_labeled SagemakerLabelingJob#max_percentage_of_input_dataset_labeled}.
	MaxPercentageOfInputDatasetLabeled *float64 `field:"optional" json:"maxPercentageOfInputDatasetLabeled" yaml:"maxPercentageOfInputDatasetLabeled"`
}

