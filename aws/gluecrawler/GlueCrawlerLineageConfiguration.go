// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerLineageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/glue_crawler#crawler_lineage_settings GlueCrawler#crawler_lineage_settings}.
	CrawlerLineageSettings *string `field:"optional" json:"crawlerLineageSettings" yaml:"crawlerLineageSettings"`
}

