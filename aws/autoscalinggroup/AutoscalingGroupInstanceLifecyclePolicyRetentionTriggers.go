// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupInstanceLifecyclePolicyRetentionTriggers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#terminate_hook_abandon AutoscalingGroup#terminate_hook_abandon}.
	TerminateHookAbandon *string `field:"optional" json:"terminateHookAbandon" yaml:"terminateHookAbandon"`
}

