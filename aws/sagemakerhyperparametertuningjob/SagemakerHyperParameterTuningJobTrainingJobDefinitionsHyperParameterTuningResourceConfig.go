// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionsHyperParameterTuningResourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#allocation_strategy SagemakerHyperParameterTuningJob#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// instance_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#instance_configs SagemakerHyperParameterTuningJob#instance_configs}
	InstanceConfigs interface{} `field:"optional" json:"instanceConfigs" yaml:"instanceConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#instance_count SagemakerHyperParameterTuningJob#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#instance_type SagemakerHyperParameterTuningJob#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#volume_kms_key_id SagemakerHyperParameterTuningJob#volume_kms_key_id}.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#volume_size_in_gb SagemakerHyperParameterTuningJob#volume_size_in_gb}.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

