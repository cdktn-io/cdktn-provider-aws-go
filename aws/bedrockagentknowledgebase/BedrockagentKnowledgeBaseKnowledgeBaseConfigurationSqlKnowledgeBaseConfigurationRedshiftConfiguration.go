// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfiguration struct {
	// query_engine_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#query_engine_configuration BedrockagentKnowledgeBase#query_engine_configuration}
	QueryEngineConfiguration interface{} `field:"optional" json:"queryEngineConfiguration" yaml:"queryEngineConfiguration"`
	// query_generation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#query_generation_configuration BedrockagentKnowledgeBase#query_generation_configuration}
	QueryGenerationConfiguration interface{} `field:"optional" json:"queryGenerationConfiguration" yaml:"queryGenerationConfiguration"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#storage_configuration BedrockagentKnowledgeBase#storage_configuration}
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

