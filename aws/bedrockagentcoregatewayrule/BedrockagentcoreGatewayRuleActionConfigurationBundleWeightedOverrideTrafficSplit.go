// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule


type BedrockagentcoreGatewayRuleActionConfigurationBundleWeightedOverrideTrafficSplit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_rule#name BedrockagentcoreGatewayRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_rule#weight BedrockagentcoreGatewayRule#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// configuration_bundle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_rule#configuration_bundle BedrockagentcoreGatewayRule#configuration_bundle}
	ConfigurationBundle interface{} `field:"optional" json:"configurationBundle" yaml:"configurationBundle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_rule#description BedrockagentcoreGatewayRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway_rule#metadata BedrockagentcoreGatewayRule#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

