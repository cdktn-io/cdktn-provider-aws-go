// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobHumanTaskConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#number_of_human_workers_per_data_object SagemakerLabelingJob#number_of_human_workers_per_data_object}.
	NumberOfHumanWorkersPerDataObject *float64 `field:"required" json:"numberOfHumanWorkersPerDataObject" yaml:"numberOfHumanWorkersPerDataObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#task_description SagemakerLabelingJob#task_description}.
	TaskDescription *string `field:"required" json:"taskDescription" yaml:"taskDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#task_time_limit_in_seconds SagemakerLabelingJob#task_time_limit_in_seconds}.
	TaskTimeLimitInSeconds *float64 `field:"required" json:"taskTimeLimitInSeconds" yaml:"taskTimeLimitInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#task_title SagemakerLabelingJob#task_title}.
	TaskTitle *string `field:"required" json:"taskTitle" yaml:"taskTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#workteam_arn SagemakerLabelingJob#workteam_arn}.
	WorkteamArn *string `field:"required" json:"workteamArn" yaml:"workteamArn"`
	// annotation_consolidation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#annotation_consolidation_config SagemakerLabelingJob#annotation_consolidation_config}
	AnnotationConsolidationConfig interface{} `field:"optional" json:"annotationConsolidationConfig" yaml:"annotationConsolidationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#max_concurrent_task_count SagemakerLabelingJob#max_concurrent_task_count}.
	MaxConcurrentTaskCount *float64 `field:"optional" json:"maxConcurrentTaskCount" yaml:"maxConcurrentTaskCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#pre_human_task_lambda_arn SagemakerLabelingJob#pre_human_task_lambda_arn}.
	PreHumanTaskLambdaArn *string `field:"optional" json:"preHumanTaskLambdaArn" yaml:"preHumanTaskLambdaArn"`
	// public_workforce_task_price block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#public_workforce_task_price SagemakerLabelingJob#public_workforce_task_price}
	PublicWorkforceTaskPrice interface{} `field:"optional" json:"publicWorkforceTaskPrice" yaml:"publicWorkforceTaskPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#task_availability_lifetime_in_seconds SagemakerLabelingJob#task_availability_lifetime_in_seconds}.
	TaskAvailabilityLifetimeInSeconds *float64 `field:"optional" json:"taskAvailabilityLifetimeInSeconds" yaml:"taskAvailabilityLifetimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#task_keywords SagemakerLabelingJob#task_keywords}.
	TaskKeywords *[]*string `field:"optional" json:"taskKeywords" yaml:"taskKeywords"`
	// ui_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_labeling_job#ui_config SagemakerLabelingJob#ui_config}
	UiConfig interface{} `field:"optional" json:"uiConfig" yaml:"uiConfig"`
}

