// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#execution_timeout_seconds BedrockagentKnowledgeBase#execution_timeout_seconds}.
	ExecutionTimeoutSeconds *float64 `field:"optional" json:"executionTimeoutSeconds" yaml:"executionTimeoutSeconds"`
	// generation_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#generation_context BedrockagentKnowledgeBase#generation_context}
	GenerationContext interface{} `field:"optional" json:"generationContext" yaml:"generationContext"`
}

