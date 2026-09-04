// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughStickinessConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#identifier BedrockagentcoreGatewayTarget#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#composite_identifier BedrockagentcoreGatewayTarget#composite_identifier}.
	CompositeIdentifier *[]*string `field:"optional" json:"compositeIdentifier" yaml:"compositeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#timeout BedrockagentcoreGatewayTarget#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

