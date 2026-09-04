// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironment struct {
	// agentcore_runtime_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_harness#agentcore_runtime_environment BedrockagentcoreHarness#agentcore_runtime_environment}
	AgentcoreRuntimeEnvironment interface{} `field:"optional" json:"agentcoreRuntimeEnvironment" yaml:"agentcoreRuntimeEnvironment"`
}

