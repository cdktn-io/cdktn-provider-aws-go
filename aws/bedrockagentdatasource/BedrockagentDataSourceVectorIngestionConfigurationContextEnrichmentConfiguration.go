// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationContextEnrichmentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/bedrockagent_data_source#type BedrockagentDataSource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// bedrock_foundation_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/bedrockagent_data_source#bedrock_foundation_model_configuration BedrockagentDataSource#bedrock_foundation_model_configuration}
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

