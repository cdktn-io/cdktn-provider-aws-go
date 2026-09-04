// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobConfigResourceLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_parallel_training_jobs SagemakerHyperParameterTuningJob#max_parallel_training_jobs}.
	MaxParallelTrainingJobs *float64 `field:"required" json:"maxParallelTrainingJobs" yaml:"maxParallelTrainingJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_number_of_training_jobs SagemakerHyperParameterTuningJob#max_number_of_training_jobs}.
	MaxNumberOfTrainingJobs *float64 `field:"optional" json:"maxNumberOfTrainingJobs" yaml:"maxNumberOfTrainingJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_runtime_in_seconds SagemakerHyperParameterTuningJob#max_runtime_in_seconds}.
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

