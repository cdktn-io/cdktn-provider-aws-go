// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationInference struct {
	// connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#connector BedrockagentcoreGatewayTarget#connector}
	Connector interface{} `field:"optional" json:"connector" yaml:"connector"`
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#provider BedrockagentcoreGatewayTarget#provider}
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
}

