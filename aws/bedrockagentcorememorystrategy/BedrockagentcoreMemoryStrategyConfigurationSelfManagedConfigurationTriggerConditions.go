// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememorystrategy


type BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditions struct {
	// message_based_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#message_based_trigger BedrockagentcoreMemoryStrategy#message_based_trigger}
	MessageBasedTrigger interface{} `field:"optional" json:"messageBasedTrigger" yaml:"messageBasedTrigger"`
	// time_based_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#time_based_trigger BedrockagentcoreMemoryStrategy#time_based_trigger}
	TimeBasedTrigger interface{} `field:"optional" json:"timeBasedTrigger" yaml:"timeBasedTrigger"`
	// token_based_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#token_based_trigger BedrockagentcoreMemoryStrategy#token_based_trigger}
	TokenBasedTrigger interface{} `field:"optional" json:"tokenBasedTrigger" yaml:"tokenBasedTrigger"`
}

