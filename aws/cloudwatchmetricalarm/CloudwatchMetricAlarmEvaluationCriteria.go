// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricalarm


type CloudwatchMetricAlarmEvaluationCriteria struct {
	// promql_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/cloudwatch_metric_alarm#promql_criteria CloudwatchMetricAlarm#promql_criteria}
	PromqlCriteria *CloudwatchMetricAlarmEvaluationCriteriaPromqlCriteria `field:"required" json:"promqlCriteria" yaml:"promqlCriteria"`
}

