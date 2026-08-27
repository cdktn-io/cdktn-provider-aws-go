// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememorystrategy


type BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#historical_context_window_size BedrockagentcoreMemoryStrategy#historical_context_window_size}.
	HistoricalContextWindowSize *float64 `field:"optional" json:"historicalContextWindowSize" yaml:"historicalContextWindowSize"`
	// invocation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#invocation_configuration BedrockagentcoreMemoryStrategy#invocation_configuration}
	InvocationConfiguration interface{} `field:"optional" json:"invocationConfiguration" yaml:"invocationConfiguration"`
	// trigger_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#trigger_conditions BedrockagentcoreMemoryStrategy#trigger_conditions}
	TriggerConditions interface{} `field:"optional" json:"triggerConditions" yaml:"triggerConditions"`
}

