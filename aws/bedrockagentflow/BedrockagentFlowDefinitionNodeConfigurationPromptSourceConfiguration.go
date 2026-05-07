// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfiguration struct {
	// inline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/bedrockagent_flow#inline BedrockagentFlow#inline}
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/bedrockagent_flow#resource BedrockagentFlow#resource}
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}

