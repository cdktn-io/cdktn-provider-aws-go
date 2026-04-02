// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionOriginGroupMember struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/cloudfront_distribution#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}

