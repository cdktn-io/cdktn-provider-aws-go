// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOriginVpcOriginConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#vpc_origin_id CloudfrontMultitenantDistribution#vpc_origin_id}.
	VpcOriginId *string `field:"required" json:"vpcOriginId" yaml:"vpcOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_keepalive_timeout CloudfrontMultitenantDistribution#origin_keepalive_timeout}.
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_read_timeout CloudfrontMultitenantDistribution#origin_read_timeout}.
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
}

