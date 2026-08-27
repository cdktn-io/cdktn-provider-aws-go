// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_pending_time_in_seconds SagemakerAlgorithm#max_pending_time_in_seconds}.
	MaxPendingTimeInSeconds *float64 `field:"optional" json:"maxPendingTimeInSeconds" yaml:"maxPendingTimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_runtime_in_seconds SagemakerAlgorithm#max_runtime_in_seconds}.
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_wait_time_in_seconds SagemakerAlgorithm#max_wait_time_in_seconds}.
	MaxWaitTimeInSeconds *float64 `field:"optional" json:"maxWaitTimeInSeconds" yaml:"maxWaitTimeInSeconds"`
}

