// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationMcpConnectorSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#connector_id BedrockagentcoreGatewayTarget#connector_id}.
	ConnectorId *string `field:"required" json:"connectorId" yaml:"connectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#version BedrockagentcoreGatewayTarget#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

