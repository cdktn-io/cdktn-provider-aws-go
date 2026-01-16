// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// provisioned_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#provisioned_configuration BedrockagentKnowledgeBase#provisioned_configuration}
	ProvisionedConfiguration interface{} `field:"optional" json:"provisionedConfiguration" yaml:"provisionedConfiguration"`
	// serverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#serverless_configuration BedrockagentKnowledgeBase#serverless_configuration}
	ServerlessConfiguration interface{} `field:"optional" json:"serverlessConfiguration" yaml:"serverlessConfiguration"`
}

