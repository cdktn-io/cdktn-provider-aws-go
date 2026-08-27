// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperExporter struct {
	// opensearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#opensearch PrometheusScraper#opensearch}
	Opensearch interface{} `field:"optional" json:"opensearch" yaml:"opensearch"`
}

