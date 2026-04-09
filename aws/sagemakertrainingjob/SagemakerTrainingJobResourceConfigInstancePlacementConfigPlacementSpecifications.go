// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobResourceConfigInstancePlacementConfigPlacementSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_training_job#instance_count SagemakerTrainingJob#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_training_job#ultra_server_id SagemakerTrainingJob#ultra_server_id}.
	UltraServerId *string `field:"optional" json:"ultraServerId" yaml:"ultraServerId"`
}

