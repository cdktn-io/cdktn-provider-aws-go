// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionRestrictions struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/cloudfront_distribution#geo_restriction CloudfrontDistribution#geo_restriction}
	GeoRestriction *CloudfrontDistributionRestrictionsGeoRestriction `field:"required" json:"geoRestriction" yaml:"geoRestriction"`
}

