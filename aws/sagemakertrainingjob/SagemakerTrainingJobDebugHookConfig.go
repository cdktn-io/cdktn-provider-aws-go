// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobDebugHookConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#s3_output_path SagemakerTrainingJob#s3_output_path}.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// collection_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#collection_configurations SagemakerTrainingJob#collection_configurations}
	CollectionConfigurations interface{} `field:"optional" json:"collectionConfigurations" yaml:"collectionConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#hook_parameters SagemakerTrainingJob#hook_parameters}.
	HookParameters *map[string]*string `field:"optional" json:"hookParameters" yaml:"hookParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#local_path SagemakerTrainingJob#local_path}.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

