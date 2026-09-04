// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricalarm


type CloudwatchMetricAlarmEvaluationCriteriaPromqlCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/cloudwatch_metric_alarm#query CloudwatchMetricAlarm#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/cloudwatch_metric_alarm#pending_period CloudwatchMetricAlarm#pending_period}.
	PendingPeriod *float64 `field:"optional" json:"pendingPeriod" yaml:"pendingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/cloudwatch_metric_alarm#recovery_period CloudwatchMetricAlarm#recovery_period}.
	RecoveryPeriod *float64 `field:"optional" json:"recoveryPeriod" yaml:"recoveryPeriod"`
}

