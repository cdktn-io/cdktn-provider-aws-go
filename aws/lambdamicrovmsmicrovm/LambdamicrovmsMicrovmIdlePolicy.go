// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmsmicrovm


type LambdamicrovmsMicrovmIdlePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/lambdamicrovms_microvm#auto_resume_enabled LambdamicrovmsMicrovm#auto_resume_enabled}.
	AutoResumeEnabled interface{} `field:"required" json:"autoResumeEnabled" yaml:"autoResumeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/lambdamicrovms_microvm#max_idle_duration_seconds LambdamicrovmsMicrovm#max_idle_duration_seconds}.
	MaxIdleDurationSeconds *float64 `field:"required" json:"maxIdleDurationSeconds" yaml:"maxIdleDurationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/lambdamicrovms_microvm#suspended_duration_seconds LambdamicrovmsMicrovm#suspended_duration_seconds}.
	SuspendedDurationSeconds *float64 `field:"required" json:"suspendedDurationSeconds" yaml:"suspendedDurationSeconds"`
}

