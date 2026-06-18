// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobRetryStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/sagemaker_training_job#maximum_retry_attempts SagemakerTrainingJob#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `field:"required" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

