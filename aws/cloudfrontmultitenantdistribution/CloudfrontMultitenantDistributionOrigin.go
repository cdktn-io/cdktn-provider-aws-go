// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#domain_name CloudfrontMultitenantDistribution#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#id CloudfrontMultitenantDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#connection_attempts CloudfrontMultitenantDistribution#connection_attempts}.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#connection_timeout CloudfrontMultitenantDistribution#connection_timeout}.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#custom_header CloudfrontMultitenantDistribution#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// custom_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#custom_origin_config CloudfrontMultitenantDistribution#custom_origin_config}
	CustomOriginConfig interface{} `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_access_control_id CloudfrontMultitenantDistribution#origin_access_control_id}.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_path CloudfrontMultitenantDistribution#origin_path}.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// origin_shield block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_shield CloudfrontMultitenantDistribution#origin_shield}
	OriginShield interface{} `field:"optional" json:"originShield" yaml:"originShield"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#response_completion_timeout CloudfrontMultitenantDistribution#response_completion_timeout}.
	ResponseCompletionTimeout *float64 `field:"optional" json:"responseCompletionTimeout" yaml:"responseCompletionTimeout"`
	// vpc_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#vpc_origin_config CloudfrontMultitenantDistribution#vpc_origin_config}
	VpcOriginConfig interface{} `field:"optional" json:"vpcOriginConfig" yaml:"vpcOriginConfig"`
}

