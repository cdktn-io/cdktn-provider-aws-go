// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememorystrategy


type BedrockagentcoreMemoryStrategyMemoryRecordSchemaMetadataSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#key BedrockagentcoreMemoryStrategy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// extraction_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#extraction_config BedrockagentcoreMemoryStrategy#extraction_config}
	ExtractionConfig interface{} `field:"optional" json:"extractionConfig" yaml:"extractionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#extraction_type BedrockagentcoreMemoryStrategy#extraction_type}.
	ExtractionType *string `field:"optional" json:"extractionType" yaml:"extractionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#type BedrockagentcoreMemoryStrategy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

