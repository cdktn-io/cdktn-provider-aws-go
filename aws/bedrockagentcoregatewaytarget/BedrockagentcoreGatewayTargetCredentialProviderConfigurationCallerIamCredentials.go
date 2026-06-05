// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetCredentialProviderConfigurationCallerIamCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/bedrockagentcore_gateway_target#service BedrockagentcoreGatewayTarget#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/bedrockagentcore_gateway_target#region BedrockagentcoreGatewayTarget#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

