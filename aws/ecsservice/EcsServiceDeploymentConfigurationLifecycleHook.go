// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationLifecycleHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#lifecycle_stages EcsService#lifecycle_stages}.
	LifecycleStages *[]*string `field:"required" json:"lifecycleStages" yaml:"lifecycleStages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#hook_details EcsService#hook_details}.
	HookDetails *string `field:"optional" json:"hookDetails" yaml:"hookDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#hook_target_arn EcsService#hook_target_arn}.
	HookTargetArn *string `field:"optional" json:"hookTargetArn" yaml:"hookTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#role_arn EcsService#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#target_type EcsService#target_type}.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// timeout_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#timeout_configuration EcsService#timeout_configuration}
	TimeoutConfiguration *EcsServiceDeploymentConfigurationLifecycleHookTimeoutConfiguration `field:"optional" json:"timeoutConfiguration" yaml:"timeoutConfiguration"`
}

