// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationInferenceProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#endpoint BedrockagentcoreGatewayTarget#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// model_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#model_mapping BedrockagentcoreGatewayTarget#model_mapping}
	ModelMapping interface{} `field:"optional" json:"modelMapping" yaml:"modelMapping"`
	// operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#operation BedrockagentcoreGatewayTarget#operation}
	Operation interface{} `field:"optional" json:"operation" yaml:"operation"`
}

