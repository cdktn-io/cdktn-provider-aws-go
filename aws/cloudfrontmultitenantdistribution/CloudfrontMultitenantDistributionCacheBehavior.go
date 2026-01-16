// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionCacheBehavior struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#path_pattern CloudfrontMultitenantDistribution#path_pattern}.
	PathPattern *string `field:"required" json:"pathPattern" yaml:"pathPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#target_origin_id CloudfrontMultitenantDistribution#target_origin_id}.
	TargetOriginId *string `field:"required" json:"targetOriginId" yaml:"targetOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#viewer_protocol_policy CloudfrontMultitenantDistribution#viewer_protocol_policy}.
	ViewerProtocolPolicy *string `field:"required" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// allowed_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#allowed_methods CloudfrontMultitenantDistribution#allowed_methods}
	AllowedMethods interface{} `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#cache_policy_id CloudfrontMultitenantDistribution#cache_policy_id}.
	CachePolicyId *string `field:"optional" json:"cachePolicyId" yaml:"cachePolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#compress CloudfrontMultitenantDistribution#compress}.
	Compress interface{} `field:"optional" json:"compress" yaml:"compress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#field_level_encryption_id CloudfrontMultitenantDistribution#field_level_encryption_id}.
	FieldLevelEncryptionId *string `field:"optional" json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#function_association CloudfrontMultitenantDistribution#function_association}
	FunctionAssociation interface{} `field:"optional" json:"functionAssociation" yaml:"functionAssociation"`
	// lambda_function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#lambda_function_association CloudfrontMultitenantDistribution#lambda_function_association}
	LambdaFunctionAssociation interface{} `field:"optional" json:"lambdaFunctionAssociation" yaml:"lambdaFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_request_policy_id CloudfrontMultitenantDistribution#origin_request_policy_id}.
	OriginRequestPolicyId *string `field:"optional" json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#realtime_log_config_arn CloudfrontMultitenantDistribution#realtime_log_config_arn}.
	RealtimeLogConfigArn *string `field:"optional" json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#response_headers_policy_id CloudfrontMultitenantDistribution#response_headers_policy_id}.
	ResponseHeadersPolicyId *string `field:"optional" json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// trusted_key_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#trusted_key_groups CloudfrontMultitenantDistribution#trusted_key_groups}
	TrustedKeyGroups interface{} `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
}

