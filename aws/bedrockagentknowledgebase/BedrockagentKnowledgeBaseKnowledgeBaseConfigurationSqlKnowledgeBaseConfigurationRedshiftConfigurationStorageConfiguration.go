// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// aws_data_catalog_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#aws_data_catalog_configuration BedrockagentKnowledgeBase#aws_data_catalog_configuration}
	AwsDataCatalogConfiguration interface{} `field:"optional" json:"awsDataCatalogConfiguration" yaml:"awsDataCatalogConfiguration"`
	// redshift_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#redshift_configuration BedrockagentKnowledgeBase#redshift_configuration}
	RedshiftConfiguration interface{} `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
}

