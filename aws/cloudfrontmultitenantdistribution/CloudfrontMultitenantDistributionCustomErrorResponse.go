// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionCustomErrorResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#error_code CloudfrontMultitenantDistribution#error_code}.
	ErrorCode *float64 `field:"required" json:"errorCode" yaml:"errorCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#error_caching_min_ttl CloudfrontMultitenantDistribution#error_caching_min_ttl}.
	ErrorCachingMinTtl *float64 `field:"optional" json:"errorCachingMinTtl" yaml:"errorCachingMinTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#response_code CloudfrontMultitenantDistribution#response_code}.
	ResponseCode *string `field:"optional" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#response_page_path CloudfrontMultitenantDistribution#response_page_path}.
	ResponsePagePath *string `field:"optional" json:"responsePagePath" yaml:"responsePagePath"`
}

