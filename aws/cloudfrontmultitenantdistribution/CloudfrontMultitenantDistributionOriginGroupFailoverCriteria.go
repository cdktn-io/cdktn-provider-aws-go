// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOriginGroupFailoverCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#status_codes CloudfrontMultitenantDistribution#status_codes}.
	StatusCodes *[]*float64 `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

