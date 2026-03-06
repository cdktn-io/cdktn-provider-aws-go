// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupTrafficSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/autoscaling_group#identifier AutoscalingGroup#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/autoscaling_group#type AutoscalingGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

