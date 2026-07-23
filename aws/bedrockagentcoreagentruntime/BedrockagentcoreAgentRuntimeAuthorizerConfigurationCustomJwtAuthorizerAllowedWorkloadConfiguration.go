// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration struct {
	// hosting_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_agent_runtime#hosting_environment BedrockagentcoreAgentRuntime#hosting_environment}
	HostingEnvironment interface{} `field:"optional" json:"hostingEnvironment" yaml:"hostingEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_agent_runtime#workload_identities BedrockagentcoreAgentRuntime#workload_identities}.
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

