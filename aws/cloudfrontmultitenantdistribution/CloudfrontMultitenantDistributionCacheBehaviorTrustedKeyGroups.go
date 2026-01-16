// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionCacheBehaviorTrustedKeyGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#enabled CloudfrontMultitenantDistribution#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#items CloudfrontMultitenantDistribution#items}.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

