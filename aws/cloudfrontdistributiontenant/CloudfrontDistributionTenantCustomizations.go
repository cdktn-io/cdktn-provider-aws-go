// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant


type CloudfrontDistributionTenantCustomizations struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#certificate CloudfrontDistributionTenant#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#geo_restriction CloudfrontDistributionTenant#geo_restriction}
	GeoRestriction interface{} `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
	// web_acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#web_acl CloudfrontDistributionTenant#web_acl}
	WebAcl interface{} `field:"optional" json:"webAcl" yaml:"webAcl"`
}

