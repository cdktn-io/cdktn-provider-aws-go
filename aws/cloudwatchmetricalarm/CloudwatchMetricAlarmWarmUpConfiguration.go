// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricalarm


type CloudwatchMetricAlarmWarmUpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/cloudwatch_metric_alarm#warm_up_period_duration_in_minutes CloudwatchMetricAlarm#warm_up_period_duration_in_minutes}.
	WarmUpPeriodDurationInMinutes *float64 `field:"required" json:"warmUpPeriodDurationInMinutes" yaml:"warmUpPeriodDurationInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/cloudwatch_metric_alarm#only_start_evaluating_after_warm_up_period_ends CloudwatchMetricAlarm#only_start_evaluating_after_warm_up_period_ends}.
	OnlyStartEvaluatingAfterWarmUpPeriodEnds interface{} `field:"optional" json:"onlyStartEvaluatingAfterWarmUpPeriodEnds" yaml:"onlyStartEvaluatingAfterWarmUpPeriodEnds"`
}

