// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOriginCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#header_name CloudfrontMultitenantDistribution#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#header_value CloudfrontMultitenantDistribution#header_value}.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

