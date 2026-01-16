// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfigurationGenerationContextCuratedQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#natural_language BedrockagentKnowledgeBase#natural_language}.
	NaturalLanguage *string `field:"required" json:"naturalLanguage" yaml:"naturalLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#sql BedrockagentKnowledgeBase#sql}.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
}

