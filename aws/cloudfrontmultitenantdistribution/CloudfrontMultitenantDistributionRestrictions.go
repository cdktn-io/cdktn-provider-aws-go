// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionRestrictions struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/cloudfront_multitenant_distribution#geo_restriction CloudfrontMultitenantDistribution#geo_restriction}
	GeoRestriction interface{} `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
}

