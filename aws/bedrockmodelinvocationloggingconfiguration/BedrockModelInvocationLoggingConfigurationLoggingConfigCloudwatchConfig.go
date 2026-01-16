// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationloggingconfiguration


type BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrock_model_invocation_logging_configuration#log_group_name BedrockModelInvocationLoggingConfiguration#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrock_model_invocation_logging_configuration#role_arn BedrockModelInvocationLoggingConfiguration#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// large_data_delivery_s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrock_model_invocation_logging_configuration#large_data_delivery_s3_config BedrockModelInvocationLoggingConfiguration#large_data_delivery_s3_config}
	LargeDataDeliveryS3Config interface{} `field:"optional" json:"largeDataDeliveryS3Config" yaml:"largeDataDeliveryS3Config"`
}

