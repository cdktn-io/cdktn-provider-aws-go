// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeAuthorizerConfiguration struct {
	// custom_jwt_authorizer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/bedrockagentcore_agent_runtime#custom_jwt_authorizer BedrockagentcoreAgentRuntime#custom_jwt_authorizer}
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

