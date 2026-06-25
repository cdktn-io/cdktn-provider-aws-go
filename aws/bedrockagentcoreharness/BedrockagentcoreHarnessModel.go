// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessModel struct {
	// bedrock_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#bedrock_model_config BedrockagentcoreHarness#bedrock_model_config}
	BedrockModelConfig interface{} `field:"optional" json:"bedrockModelConfig" yaml:"bedrockModelConfig"`
	// gemini_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#gemini_model_config BedrockagentcoreHarness#gemini_model_config}
	GeminiModelConfig interface{} `field:"optional" json:"geminiModelConfig" yaml:"geminiModelConfig"`
	// openai_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#openai_model_config BedrockagentcoreHarness#openai_model_config}
	OpenaiModelConfig interface{} `field:"optional" json:"openaiModelConfig" yaml:"openaiModelConfig"`
}

