// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseStorageConfigurationMongoDbAtlasConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#collection_name BedrockagentKnowledgeBase#collection_name}.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#credentials_secret_arn BedrockagentKnowledgeBase#credentials_secret_arn}.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#database_name BedrockagentKnowledgeBase#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#endpoint BedrockagentKnowledgeBase#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#vector_index_name BedrockagentKnowledgeBase#vector_index_name}.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#endpoint_service_name BedrockagentKnowledgeBase#endpoint_service_name}.
	EndpointServiceName *string `field:"optional" json:"endpointServiceName" yaml:"endpointServiceName"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#field_mapping BedrockagentKnowledgeBase#field_mapping}
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#text_index_name BedrockagentKnowledgeBase#text_index_name}.
	TextIndexName *string `field:"optional" json:"textIndexName" yaml:"textIndexName"`
}

