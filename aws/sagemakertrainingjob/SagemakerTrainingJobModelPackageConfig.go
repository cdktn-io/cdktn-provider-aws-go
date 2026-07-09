// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobModelPackageConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_training_job#model_package_group_arn SagemakerTrainingJob#model_package_group_arn}.
	ModelPackageGroupArn *string `field:"required" json:"modelPackageGroupArn" yaml:"modelPackageGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_training_job#source_model_package_arn SagemakerTrainingJob#source_model_package_arn}.
	SourceModelPackageArn *string `field:"optional" json:"sourceModelPackageArn" yaml:"sourceModelPackageArn"`
}

