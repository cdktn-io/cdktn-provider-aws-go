// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpApiGatewayApiGatewayToolConfigurationToolFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/bedrockagentcore_gateway_target#filter_path BedrockagentcoreGatewayTarget#filter_path}.
	FilterPath *string `field:"required" json:"filterPath" yaml:"filterPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/bedrockagentcore_gateway_target#methods BedrockagentcoreGatewayTarget#methods}.
	Methods *[]*string `field:"required" json:"methods" yaml:"methods"`
}

