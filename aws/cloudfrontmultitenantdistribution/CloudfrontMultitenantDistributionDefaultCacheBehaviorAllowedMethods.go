// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionDefaultCacheBehaviorAllowedMethods struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#cached_methods CloudfrontMultitenantDistribution#cached_methods}.
	CachedMethods *[]*string `field:"required" json:"cachedMethods" yaml:"cachedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#items CloudfrontMultitenantDistribution#items}.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
}

