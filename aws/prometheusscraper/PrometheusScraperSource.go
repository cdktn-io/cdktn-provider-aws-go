// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperSource struct {
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/prometheus_scraper#eks PrometheusScraper#eks}
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/prometheus_scraper#vpc PrometheusScraper#vpc}
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

