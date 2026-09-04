// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupInstanceLifecyclePolicy struct {
	// retention_triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/autoscaling_group#retention_triggers AutoscalingGroup#retention_triggers}
	RetentionTriggers *AutoscalingGroupInstanceLifecyclePolicyRetentionTriggers `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}

