// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessMemory struct {
	// agentcore_memory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/bedrockagentcore_harness#agentcore_memory_configuration BedrockagentcoreHarness#agentcore_memory_configuration}
	AgentcoreMemoryConfiguration interface{} `field:"optional" json:"agentcoreMemoryConfiguration" yaml:"agentcoreMemoryConfiguration"`
}

