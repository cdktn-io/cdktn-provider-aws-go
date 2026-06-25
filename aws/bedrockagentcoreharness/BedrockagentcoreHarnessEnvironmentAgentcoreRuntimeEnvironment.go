// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentAgentcoreRuntimeEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#agent_runtime_arn BedrockagentcoreHarness#agent_runtime_arn}.
	AgentRuntimeArn *string `field:"optional" json:"agentRuntimeArn" yaml:"agentRuntimeArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#agent_runtime_id BedrockagentcoreHarness#agent_runtime_id}.
	AgentRuntimeId *string `field:"optional" json:"agentRuntimeId" yaml:"agentRuntimeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#agent_runtime_name BedrockagentcoreHarness#agent_runtime_name}.
	AgentRuntimeName *string `field:"optional" json:"agentRuntimeName" yaml:"agentRuntimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#filesystem_configuration BedrockagentcoreHarness#filesystem_configuration}.
	FilesystemConfiguration interface{} `field:"optional" json:"filesystemConfiguration" yaml:"filesystemConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#lifecycle_configuration BedrockagentcoreHarness#lifecycle_configuration}.
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#network_configuration BedrockagentcoreHarness#network_configuration}.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
}

