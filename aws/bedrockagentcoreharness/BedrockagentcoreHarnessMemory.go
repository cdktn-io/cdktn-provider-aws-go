// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessMemory struct {
	// agentcore_memory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#agentcore_memory_configuration BedrockagentcoreHarness#agentcore_memory_configuration}
	AgentcoreMemoryConfiguration interface{} `field:"optional" json:"agentcoreMemoryConfiguration" yaml:"agentcoreMemoryConfiguration"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#disabled BedrockagentcoreHarness#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// managed_memory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#managed_memory_configuration BedrockagentcoreHarness#managed_memory_configuration}
	ManagedMemoryConfiguration interface{} `field:"optional" json:"managedMemoryConfiguration" yaml:"managedMemoryConfiguration"`
}

