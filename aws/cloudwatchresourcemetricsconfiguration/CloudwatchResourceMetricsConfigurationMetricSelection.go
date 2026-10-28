// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchresourcemetricsconfiguration


type CloudwatchResourceMetricsConfigurationMetricSelection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/cloudwatch_resource_metrics_configuration#include_metrics CloudwatchResourceMetricsConfiguration#include_metrics}.
	IncludeMetrics *[]*string `field:"required" json:"includeMetrics" yaml:"includeMetrics"`
}

