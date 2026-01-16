// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// redshift_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#redshift_configuration BedrockagentKnowledgeBase#redshift_configuration}
	RedshiftConfiguration interface{} `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
}

