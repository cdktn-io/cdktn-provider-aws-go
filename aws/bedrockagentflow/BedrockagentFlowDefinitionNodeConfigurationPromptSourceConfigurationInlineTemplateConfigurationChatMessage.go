// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfigurationChatMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagent_flow#role BedrockagentFlow#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagent_flow#content BedrockagentFlow#content}
	Content interface{} `field:"optional" json:"content" yaml:"content"`
}

