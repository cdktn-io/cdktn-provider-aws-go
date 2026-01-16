// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfigurationProvisionedConfigurationAuthConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#database_user BedrockagentKnowledgeBase#database_user}.
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#username_password_secret_arn BedrockagentKnowledgeBase#username_password_secret_arn}.
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

