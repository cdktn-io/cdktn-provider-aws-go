// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobInputConfigDataSourceSnsDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_labeling_job#sns_topic_arn SagemakerLabelingJob#sns_topic_arn}.
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

