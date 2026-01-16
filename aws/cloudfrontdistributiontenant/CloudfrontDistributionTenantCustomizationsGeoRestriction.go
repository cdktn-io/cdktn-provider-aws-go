// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant


type CloudfrontDistributionTenantCustomizationsGeoRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#locations CloudfrontDistributionTenant#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#restriction_type CloudfrontDistributionTenant#restriction_type}.
	RestrictionType *string `field:"optional" json:"restrictionType" yaml:"restrictionType"`
}

