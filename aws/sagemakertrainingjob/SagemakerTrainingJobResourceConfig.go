// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobResourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#instance_count SagemakerTrainingJob#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// instance_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#instance_groups SagemakerTrainingJob#instance_groups}
	InstanceGroups interface{} `field:"optional" json:"instanceGroups" yaml:"instanceGroups"`
	// instance_placement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#instance_placement_config SagemakerTrainingJob#instance_placement_config}
	InstancePlacementConfig interface{} `field:"optional" json:"instancePlacementConfig" yaml:"instancePlacementConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#instance_type SagemakerTrainingJob#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#keep_alive_period_in_seconds SagemakerTrainingJob#keep_alive_period_in_seconds}.
	KeepAlivePeriodInSeconds *float64 `field:"optional" json:"keepAlivePeriodInSeconds" yaml:"keepAlivePeriodInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#training_plan_arn SagemakerTrainingJob#training_plan_arn}.
	TrainingPlanArn *string `field:"optional" json:"trainingPlanArn" yaml:"trainingPlanArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#volume_kms_key_id SagemakerTrainingJob#volume_kms_key_id}.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_training_job#volume_size_in_gb SagemakerTrainingJob#volume_size_in_gb}.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

