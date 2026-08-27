// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOriginGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#id CloudfrontMultitenantDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// failover_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#failover_criteria CloudfrontMultitenantDistribution#failover_criteria}
	FailoverCriteria interface{} `field:"optional" json:"failoverCriteria" yaml:"failoverCriteria"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#member CloudfrontMultitenantDistribution#member}
	Member interface{} `field:"optional" json:"member" yaml:"member"`
}

