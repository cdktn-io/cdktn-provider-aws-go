// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpMcpServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#endpoint BedrockagentcoreGatewayTarget#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#listing_mode BedrockagentcoreGatewayTarget#listing_mode}.
	ListingMode *string `field:"optional" json:"listingMode" yaml:"listingMode"`
}

