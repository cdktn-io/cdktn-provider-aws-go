// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_data_source#connector_parameters BedrockagentDataSource#connector_parameters}.
	ConnectorParameters *string `field:"optional" json:"connectorParameters" yaml:"connectorParameters"`
	// deletion_protection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_data_source#deletion_protection_configuration BedrockagentDataSource#deletion_protection_configuration}
	DeletionProtectionConfiguration interface{} `field:"optional" json:"deletionProtectionConfiguration" yaml:"deletionProtectionConfiguration"`
	// media_extraction_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagent_data_source#media_extraction_configuration BedrockagentDataSource#media_extraction_configuration}
	MediaExtractionConfiguration interface{} `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
}

