// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessTool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#type BedrockagentcoreHarness#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#config BedrockagentcoreHarness#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#name BedrockagentcoreHarness#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

