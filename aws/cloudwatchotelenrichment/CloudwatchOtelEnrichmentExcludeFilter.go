// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchotelenrichment


type CloudwatchOtelEnrichmentExcludeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/cloudwatch_otel_enrichment#namespace CloudwatchOtelEnrichment#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/cloudwatch_otel_enrichment#metric_names CloudwatchOtelEnrichment#metric_names}.
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

