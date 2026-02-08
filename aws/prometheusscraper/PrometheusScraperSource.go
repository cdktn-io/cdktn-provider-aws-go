// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperSource struct {
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/prometheus_scraper#eks PrometheusScraper#eks}
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
}

