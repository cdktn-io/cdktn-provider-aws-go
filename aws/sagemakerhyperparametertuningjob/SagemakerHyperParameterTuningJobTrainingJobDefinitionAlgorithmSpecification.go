// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionAlgorithmSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_input_mode SagemakerHyperParameterTuningJob#training_input_mode}.
	TrainingInputMode *string `field:"required" json:"trainingInputMode" yaml:"trainingInputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#algorithm_name SagemakerHyperParameterTuningJob#algorithm_name}.
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// metric_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#metric_definitions SagemakerHyperParameterTuningJob#metric_definitions}
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_image SagemakerHyperParameterTuningJob#training_image}.
	TrainingImage *string `field:"optional" json:"trainingImage" yaml:"trainingImage"`
}

