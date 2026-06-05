// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionResourceConfigInstanceGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_algorithm#instance_count SagemakerAlgorithm#instance_count}.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_algorithm#instance_group_name SagemakerAlgorithm#instance_group_name}.
	InstanceGroupName *string `field:"required" json:"instanceGroupName" yaml:"instanceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_algorithm#instance_type SagemakerAlgorithm#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

