// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionDefaultCacheBehaviorLambdaFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#event_type CloudfrontMultitenantDistribution#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#lambda_function_arn CloudfrontMultitenantDistribution#lambda_function_arn}.
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#include_body CloudfrontMultitenantDistribution#include_body}.
	IncludeBody interface{} `field:"optional" json:"includeBody" yaml:"includeBody"`
}

