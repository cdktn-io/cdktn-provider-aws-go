// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagentcore_gateway_target#lambda_arn BedrockagentcoreGatewayTarget#lambda_arn}.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// tool_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagentcore_gateway_target#tool_schema BedrockagentcoreGatewayTarget#tool_schema}
	ToolSchema interface{} `field:"optional" json:"toolSchema" yaml:"toolSchema"`
}

