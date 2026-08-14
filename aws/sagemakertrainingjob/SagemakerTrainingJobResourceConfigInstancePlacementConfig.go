// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobResourceConfigInstancePlacementConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_training_job#enable_multiple_jobs SagemakerTrainingJob#enable_multiple_jobs}.
	EnableMultipleJobs interface{} `field:"optional" json:"enableMultipleJobs" yaml:"enableMultipleJobs"`
	// placement_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_training_job#placement_specifications SagemakerTrainingJob#placement_specifications}
	PlacementSpecifications interface{} `field:"optional" json:"placementSpecifications" yaml:"placementSpecifications"`
}

