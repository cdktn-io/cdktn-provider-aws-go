// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessMemoryAgentcoreMemoryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/bedrockagentcore_harness#arn BedrockagentcoreHarness#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/bedrockagentcore_harness#actor_id BedrockagentcoreHarness#actor_id}.
	ActorId *string `field:"optional" json:"actorId" yaml:"actorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/bedrockagentcore_harness#messages_count BedrockagentcoreHarness#messages_count}.
	MessagesCount *float64 `field:"optional" json:"messagesCount" yaml:"messagesCount"`
	// retrieval_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/bedrockagentcore_harness#retrieval_config BedrockagentcoreHarness#retrieval_config}
	RetrievalConfig interface{} `field:"optional" json:"retrievalConfig" yaml:"retrievalConfig"`
}

