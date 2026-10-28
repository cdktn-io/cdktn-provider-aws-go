// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricalarm


type CloudwatchMetricAlarmEvaluationWindow struct {
	// wall_clock_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/cloudwatch_metric_alarm#wall_clock_window CloudwatchMetricAlarm#wall_clock_window}
	WallClockWindow *CloudwatchMetricAlarmEvaluationWindowWallClockWindow `field:"required" json:"wallClockWindow" yaml:"wallClockWindow"`
}

