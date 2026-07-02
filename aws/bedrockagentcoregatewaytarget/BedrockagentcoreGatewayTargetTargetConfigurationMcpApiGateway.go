// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpApiGateway struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_gateway_target#rest_api_id BedrockagentcoreGatewayTarget#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_gateway_target#stage BedrockagentcoreGatewayTarget#stage}.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// api_gateway_tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_gateway_target#api_gateway_tool_configuration BedrockagentcoreGatewayTarget#api_gateway_tool_configuration}
	ApiGatewayToolConfiguration interface{} `field:"optional" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
}

