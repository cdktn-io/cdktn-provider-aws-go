// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionRestrictionsGeoRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#restriction_type CloudfrontMultitenantDistribution#restriction_type}.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#items CloudfrontMultitenantDistribution#items}.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

