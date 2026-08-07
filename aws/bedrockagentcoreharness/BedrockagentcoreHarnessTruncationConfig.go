// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessTruncationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagentcore_harness#sliding_window BedrockagentcoreHarness#sliding_window}.
	SlidingWindow interface{} `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagentcore_harness#summarization BedrockagentcoreHarness#summarization}.
	Summarization interface{} `field:"optional" json:"summarization" yaml:"summarization"`
}

