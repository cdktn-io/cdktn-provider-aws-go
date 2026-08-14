// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPrompt struct {
	// guardrail_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagent_flow#guardrail_configuration BedrockagentFlow#guardrail_configuration}
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagent_flow#source_configuration BedrockagentFlow#source_configuration}
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

