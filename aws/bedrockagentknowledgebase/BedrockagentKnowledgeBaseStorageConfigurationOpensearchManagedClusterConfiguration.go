// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#domain_arn BedrockagentKnowledgeBase#domain_arn}.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#domain_endpoint BedrockagentKnowledgeBase#domain_endpoint}.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#vector_index_name BedrockagentKnowledgeBase#vector_index_name}.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/bedrockagent_knowledge_base#field_mapping BedrockagentKnowledgeBase#field_mapping}
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
}

