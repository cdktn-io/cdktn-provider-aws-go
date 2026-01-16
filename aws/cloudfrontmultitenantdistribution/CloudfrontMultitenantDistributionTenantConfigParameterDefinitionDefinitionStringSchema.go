// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionTenantConfigParameterDefinitionDefinitionStringSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#required CloudfrontMultitenantDistribution#required}.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#comment CloudfrontMultitenantDistribution#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#default_value CloudfrontMultitenantDistribution#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}

