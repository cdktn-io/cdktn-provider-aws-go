// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationParsingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/bedrockagent_data_source#parsing_strategy BedrockagentDataSource#parsing_strategy}.
	ParsingStrategy *string `field:"required" json:"parsingStrategy" yaml:"parsingStrategy"`
	// bedrock_data_automation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/bedrockagent_data_source#bedrock_data_automation_configuration BedrockagentDataSource#bedrock_data_automation_configuration}
	BedrockDataAutomationConfiguration interface{} `field:"optional" json:"bedrockDataAutomationConfiguration" yaml:"bedrockDataAutomationConfiguration"`
	// bedrock_foundation_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/bedrockagent_data_source#bedrock_foundation_model_configuration BedrockagentDataSource#bedrock_foundation_model_configuration}
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

