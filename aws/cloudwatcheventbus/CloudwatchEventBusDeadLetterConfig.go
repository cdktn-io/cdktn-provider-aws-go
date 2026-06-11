// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventbus


type CloudwatchEventBusDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/cloudwatch_event_bus#arn CloudwatchEventBus#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

