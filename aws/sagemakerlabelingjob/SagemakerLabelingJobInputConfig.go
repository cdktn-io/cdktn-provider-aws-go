// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobInputConfig struct {
	// data_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_labeling_job#data_attributes SagemakerLabelingJob#data_attributes}
	DataAttributes interface{} `field:"optional" json:"dataAttributes" yaml:"dataAttributes"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_labeling_job#data_source SagemakerLabelingJob#data_source}
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
}

