// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionTenantConfigParameterDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#name CloudfrontMultitenantDistribution#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#definition CloudfrontMultitenantDistribution#definition}
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
}

