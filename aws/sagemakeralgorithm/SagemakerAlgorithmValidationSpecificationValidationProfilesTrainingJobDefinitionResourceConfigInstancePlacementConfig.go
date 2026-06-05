// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionResourceConfigInstancePlacementConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_algorithm#enable_multiple_jobs SagemakerAlgorithm#enable_multiple_jobs}.
	EnableMultipleJobs interface{} `field:"optional" json:"enableMultipleJobs" yaml:"enableMultipleJobs"`
	// placement_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_algorithm#placement_specifications SagemakerAlgorithm#placement_specifications}
	PlacementSpecifications interface{} `field:"optional" json:"placementSpecifications" yaml:"placementSpecifications"`
}

