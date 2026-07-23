// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#profile_name SagemakerAlgorithm#profile_name}.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// training_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#training_job_definition SagemakerAlgorithm#training_job_definition}
	TrainingJobDefinition interface{} `field:"optional" json:"trainingJobDefinition" yaml:"trainingJobDefinition"`
	// transform_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#transform_job_definition SagemakerAlgorithm#transform_job_definition}
	TransformJobDefinition interface{} `field:"optional" json:"transformJobDefinition" yaml:"transformJobDefinition"`
}

