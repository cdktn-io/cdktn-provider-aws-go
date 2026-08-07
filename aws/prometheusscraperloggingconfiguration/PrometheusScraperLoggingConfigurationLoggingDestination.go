// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusscraperloggingconfiguration


type PrometheusScraperLoggingConfigurationLoggingDestination struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/prometheus_scraper_logging_configuration#cloudwatch_logs PrometheusScraperLoggingConfiguration#cloudwatch_logs}
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
}

