// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration struct {
	// audio_extraction_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagent_data_source#audio_extraction_configuration BedrockagentDataSource#audio_extraction_configuration}
	AudioExtractionConfiguration interface{} `field:"optional" json:"audioExtractionConfiguration" yaml:"audioExtractionConfiguration"`
	// image_extraction_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagent_data_source#image_extraction_configuration BedrockagentDataSource#image_extraction_configuration}
	ImageExtractionConfiguration interface{} `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
	// video_extraction_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagent_data_source#video_extraction_configuration BedrockagentDataSource#video_extraction_configuration}
	VideoExtractionConfiguration interface{} `field:"optional" json:"videoExtractionConfiguration" yaml:"videoExtractionConfiguration"`
}

