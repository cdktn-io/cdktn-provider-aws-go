// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobHumanTaskConfigUiConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_labeling_job#human_task_ui_arn SagemakerLabelingJob#human_task_ui_arn}.
	HumanTaskUiArn *string `field:"optional" json:"humanTaskUiArn" yaml:"humanTaskUiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_labeling_job#ui_template_s3_uri SagemakerLabelingJob#ui_template_s3_uri}.
	UiTemplateS3Uri *string `field:"optional" json:"uiTemplateS3Uri" yaml:"uiTemplateS3Uri"`
}

