// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_agent_runtime#domain BedrockagentcoreAgentRuntime#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_agent_runtime#private_endpoint BedrockagentcoreAgentRuntime#private_endpoint}
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

