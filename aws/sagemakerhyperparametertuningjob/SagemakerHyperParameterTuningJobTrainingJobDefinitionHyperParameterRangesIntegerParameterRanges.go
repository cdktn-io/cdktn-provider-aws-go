// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesIntegerParameterRanges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_value SagemakerHyperParameterTuningJob#max_value}.
	MaxValue *string `field:"required" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_hyper_parameter_tuning_job#min_value SagemakerHyperParameterTuningJob#min_value}.
	MinValue *string `field:"required" json:"minValue" yaml:"minValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_hyper_parameter_tuning_job#name SagemakerHyperParameterTuningJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_hyper_parameter_tuning_job#scaling_type SagemakerHyperParameterTuningJob#scaling_type}.
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
}

