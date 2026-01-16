// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfigurationAwsDataCatalogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#table_names BedrockagentKnowledgeBase#table_names}.
	TableNames *[]*string `field:"required" json:"tableNames" yaml:"tableNames"`
}

