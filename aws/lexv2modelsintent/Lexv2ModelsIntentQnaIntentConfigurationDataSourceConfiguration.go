// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfiguration struct {
	// bedrock_knowledge_store_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/lexv2models_intent#bedrock_knowledge_store_configuration Lexv2ModelsIntent#bedrock_knowledge_store_configuration}
	BedrockKnowledgeStoreConfiguration interface{} `field:"optional" json:"bedrockKnowledgeStoreConfiguration" yaml:"bedrockKnowledgeStoreConfiguration"`
	// kendra_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/lexv2models_intent#kendra_configuration Lexv2ModelsIntent#kendra_configuration}
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// opensearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/lexv2models_intent#opensearch_configuration Lexv2ModelsIntent#opensearch_configuration}
	OpensearchConfiguration interface{} `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
}

