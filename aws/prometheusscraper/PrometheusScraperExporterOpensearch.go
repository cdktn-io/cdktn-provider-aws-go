// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperExporterOpensearch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/prometheus_scraper#domain_arn PrometheusScraper#domain_arn}.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
}

