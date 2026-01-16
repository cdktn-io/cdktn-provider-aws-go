// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOriginGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_id CloudfrontMultitenantDistribution#origin_id}.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
	// failover_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#failover_criteria CloudfrontMultitenantDistribution#failover_criteria}
	FailoverCriteria interface{} `field:"optional" json:"failoverCriteria" yaml:"failoverCriteria"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#member CloudfrontMultitenantDistribution#member}
	Member interface{} `field:"optional" json:"member" yaml:"member"`
}

