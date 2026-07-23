// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/lexv2models_intent#domain_endpoint Lexv2ModelsIntent#domain_endpoint}.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/lexv2models_intent#index_name Lexv2ModelsIntent#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/lexv2models_intent#exact_response Lexv2ModelsIntent#exact_response}.
	ExactResponse interface{} `field:"optional" json:"exactResponse" yaml:"exactResponse"`
	// exact_response_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/lexv2models_intent#exact_response_fields Lexv2ModelsIntent#exact_response_fields}
	ExactResponseFields interface{} `field:"optional" json:"exactResponseFields" yaml:"exactResponseFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/lexv2models_intent#include_fields Lexv2ModelsIntent#include_fields}.
	IncludeFields *[]*string `field:"optional" json:"includeFields" yaml:"includeFields"`
}

