// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfigurationGenerationContext struct {
	// curated_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#curated_query BedrockagentKnowledgeBase#curated_query}
	CuratedQuery interface{} `field:"optional" json:"curatedQuery" yaml:"curatedQuery"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#table BedrockagentKnowledgeBase#table}
	Table interface{} `field:"optional" json:"table" yaml:"table"`
}

