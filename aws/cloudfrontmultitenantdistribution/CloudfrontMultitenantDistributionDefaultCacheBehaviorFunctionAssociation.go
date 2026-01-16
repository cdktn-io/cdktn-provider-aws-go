// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionDefaultCacheBehaviorFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#event_type CloudfrontMultitenantDistribution#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#function_arn CloudfrontMultitenantDistribution#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

