// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolConfigInlineFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#description BedrockagentcoreHarness#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#input_schema BedrockagentcoreHarness#input_schema}.
	InputSchema *string `field:"required" json:"inputSchema" yaml:"inputSchema"`
}

