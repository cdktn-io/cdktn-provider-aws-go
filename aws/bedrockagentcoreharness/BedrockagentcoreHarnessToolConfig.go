// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolConfig struct {
	// agentcore_browser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/bedrockagentcore_harness#agentcore_browser BedrockagentcoreHarness#agentcore_browser}
	AgentcoreBrowser interface{} `field:"optional" json:"agentcoreBrowser" yaml:"agentcoreBrowser"`
	// agentcore_code_interpreter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/bedrockagentcore_harness#agentcore_code_interpreter BedrockagentcoreHarness#agentcore_code_interpreter}
	AgentcoreCodeInterpreter interface{} `field:"optional" json:"agentcoreCodeInterpreter" yaml:"agentcoreCodeInterpreter"`
	// agentcore_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/bedrockagentcore_harness#agentcore_gateway BedrockagentcoreHarness#agentcore_gateway}
	AgentcoreGateway interface{} `field:"optional" json:"agentcoreGateway" yaml:"agentcoreGateway"`
	// inline_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/bedrockagentcore_harness#inline_function BedrockagentcoreHarness#inline_function}
	InlineFunction interface{} `field:"optional" json:"inlineFunction" yaml:"inlineFunction"`
	// remote_mcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/bedrockagentcore_harness#remote_mcp BedrockagentcoreHarness#remote_mcp}
	RemoteMcp interface{} `field:"optional" json:"remoteMcp" yaml:"remoteMcp"`
}

