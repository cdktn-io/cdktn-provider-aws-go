// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationLifecycleHookTimeoutConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#action EcsService#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/ecs_service#timeout_in_minutes EcsService#timeout_in_minutes}.
	TimeoutInMinutes *string `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

