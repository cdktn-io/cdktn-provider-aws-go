// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobAlgorithmSpecificationTrainingImageConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/sagemaker_training_job#training_repository_access_mode SagemakerTrainingJob#training_repository_access_mode}.
	TrainingRepositoryAccessMode *string `field:"optional" json:"trainingRepositoryAccessMode" yaml:"trainingRepositoryAccessMode"`
	// training_repository_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/sagemaker_training_job#training_repository_auth_config SagemakerTrainingJob#training_repository_auth_config}
	TrainingRepositoryAuthConfig interface{} `field:"optional" json:"trainingRepositoryAuthConfig" yaml:"trainingRepositoryAuthConfig"`
}

