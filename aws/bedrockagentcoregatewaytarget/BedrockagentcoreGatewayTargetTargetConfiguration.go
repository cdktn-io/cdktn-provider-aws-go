// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfiguration struct {
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#http BedrockagentcoreGatewayTarget#http}
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#inference BedrockagentcoreGatewayTarget#inference}
	Inference interface{} `field:"optional" json:"inference" yaml:"inference"`
	// mcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#mcp BedrockagentcoreGatewayTarget#mcp}
	Mcp interface{} `field:"optional" json:"mcp" yaml:"mcp"`
}

