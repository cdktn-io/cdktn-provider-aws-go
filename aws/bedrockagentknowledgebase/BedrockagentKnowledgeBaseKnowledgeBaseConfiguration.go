// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// kendra_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#kendra_knowledge_base_configuration BedrockagentKnowledgeBase#kendra_knowledge_base_configuration}
	KendraKnowledgeBaseConfiguration interface{} `field:"optional" json:"kendraKnowledgeBaseConfiguration" yaml:"kendraKnowledgeBaseConfiguration"`
	// sql_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#sql_knowledge_base_configuration BedrockagentKnowledgeBase#sql_knowledge_base_configuration}
	SqlKnowledgeBaseConfiguration interface{} `field:"optional" json:"sqlKnowledgeBaseConfiguration" yaml:"sqlKnowledgeBaseConfiguration"`
	// vector_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#vector_knowledge_base_configuration BedrockagentKnowledgeBase#vector_knowledge_base_configuration}
	VectorKnowledgeBaseConfiguration interface{} `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

