// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationBedrockKnowledgeStoreConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/lexv2models_intent#bedrock_knowledge_base_arn Lexv2ModelsIntent#bedrock_knowledge_base_arn}.
	BedrockKnowledgeBaseArn *string `field:"required" json:"bedrockKnowledgeBaseArn" yaml:"bedrockKnowledgeBaseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/lexv2models_intent#exact_response Lexv2ModelsIntent#exact_response}.
	ExactResponse interface{} `field:"optional" json:"exactResponse" yaml:"exactResponse"`
	// exact_response_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/lexv2models_intent#exact_response_fields Lexv2ModelsIntent#exact_response_fields}
	ExactResponseFields interface{} `field:"optional" json:"exactResponseFields" yaml:"exactResponseFields"`
}

