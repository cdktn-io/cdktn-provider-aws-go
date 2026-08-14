// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpMcpServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_target#endpoint BedrockagentcoreGatewayTarget#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_target#listing_mode BedrockagentcoreGatewayTarget#listing_mode}.
	ListingMode *string `field:"optional" json:"listingMode" yaml:"listingMode"`
	// mcp_tool_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_target#mcp_tool_schema BedrockagentcoreGatewayTarget#mcp_tool_schema}
	McpToolSchema interface{} `field:"optional" json:"mcpToolSchema" yaml:"mcpToolSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_target#resource_priority BedrockagentcoreGatewayTarget#resource_priority}.
	ResourcePriority *float64 `field:"optional" json:"resourcePriority" yaml:"resourcePriority"`
}

