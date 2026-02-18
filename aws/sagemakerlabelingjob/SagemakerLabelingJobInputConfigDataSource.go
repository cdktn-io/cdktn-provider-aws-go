// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobInputConfigDataSource struct {
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/sagemaker_labeling_job#s3_data_source SagemakerLabelingJob#s3_data_source}
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
	// sns_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/sagemaker_labeling_job#sns_data_source SagemakerLabelingJob#sns_data_source}
	SnsDataSource interface{} `field:"optional" json:"snsDataSource" yaml:"snsDataSource"`
}

