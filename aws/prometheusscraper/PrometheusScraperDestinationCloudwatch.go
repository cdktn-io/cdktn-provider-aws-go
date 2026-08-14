// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperDestinationCloudwatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/prometheus_scraper#dataset_arn PrometheusScraper#dataset_arn}.
	DatasetArn *string `field:"required" json:"datasetArn" yaml:"datasetArn"`
}

