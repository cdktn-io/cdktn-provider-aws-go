// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRanges struct {
	// auto_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#auto_parameters SagemakerHyperParameterTuningJob#auto_parameters}
	AutoParameters interface{} `field:"optional" json:"autoParameters" yaml:"autoParameters"`
	// categorical_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#categorical_parameter_ranges SagemakerHyperParameterTuningJob#categorical_parameter_ranges}
	CategoricalParameterRanges interface{} `field:"optional" json:"categoricalParameterRanges" yaml:"categoricalParameterRanges"`
	// continuous_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#continuous_parameter_ranges SagemakerHyperParameterTuningJob#continuous_parameter_ranges}
	ContinuousParameterRanges interface{} `field:"optional" json:"continuousParameterRanges" yaml:"continuousParameterRanges"`
	// integer_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#integer_parameter_ranges SagemakerHyperParameterTuningJob#integer_parameter_ranges}
	IntegerParameterRanges interface{} `field:"optional" json:"integerParameterRanges" yaml:"integerParameterRanges"`
}

