// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobResourceConfigInstanceGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_training_job#instance_count SagemakerTrainingJob#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_training_job#instance_group_name SagemakerTrainingJob#instance_group_name}.
	InstanceGroupName *string `field:"optional" json:"instanceGroupName" yaml:"instanceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_training_job#instance_type SagemakerTrainingJob#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
}

