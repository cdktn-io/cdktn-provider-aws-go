// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobLabelingJobAlgorithmsConfigLabelingJobResourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/sagemaker_labeling_job#volume_kms_key_id SagemakerLabelingJob#volume_kms_key_id}.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/sagemaker_labeling_job#vpc_config SagemakerLabelingJob#vpc_config}
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

