// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfigurationServerlessConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#workgroup_arn BedrockagentKnowledgeBase#workgroup_arn}.
	WorkgroupArn *string `field:"required" json:"workgroupArn" yaml:"workgroupArn"`
	// auth_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#auth_configuration BedrockagentKnowledgeBase#auth_configuration}
	AuthConfiguration interface{} `field:"optional" json:"authConfiguration" yaml:"authConfiguration"`
}

