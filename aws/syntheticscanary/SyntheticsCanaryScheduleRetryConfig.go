// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryScheduleRetryConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/synthetics_canary#max_retries SyntheticsCanary#max_retries}.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
}

