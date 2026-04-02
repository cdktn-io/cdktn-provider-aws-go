// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentQnaIntentConfiguration struct {
	// bedrock_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/lexv2models_intent#bedrock_model_configuration Lexv2ModelsIntent#bedrock_model_configuration}
	BedrockModelConfiguration interface{} `field:"optional" json:"bedrockModelConfiguration" yaml:"bedrockModelConfiguration"`
	// data_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/lexv2models_intent#data_source_configuration Lexv2ModelsIntent#data_source_configuration}
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
}

