// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionStoppingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_pending_time_in_seconds SagemakerHyperParameterTuningJob#max_pending_time_in_seconds}.
	MaxPendingTimeInSeconds *float64 `field:"optional" json:"maxPendingTimeInSeconds" yaml:"maxPendingTimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_runtime_in_seconds SagemakerHyperParameterTuningJob#max_runtime_in_seconds}.
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_wait_time_in_seconds SagemakerHyperParameterTuningJob#max_wait_time_in_seconds}.
	MaxWaitTimeInSeconds *float64 `field:"optional" json:"maxWaitTimeInSeconds" yaml:"maxWaitTimeInSeconds"`
}

