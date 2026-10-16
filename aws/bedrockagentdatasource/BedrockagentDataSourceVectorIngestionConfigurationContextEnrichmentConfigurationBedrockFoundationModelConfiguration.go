// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationContextEnrichmentConfigurationBedrockFoundationModelConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/bedrockagent_data_source#model_arn BedrockagentDataSource#model_arn}.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// enrichment_strategy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/bedrockagent_data_source#enrichment_strategy_configuration BedrockagentDataSource#enrichment_strategy_configuration}
	EnrichmentStrategyConfiguration interface{} `field:"optional" json:"enrichmentStrategyConfiguration" yaml:"enrichmentStrategyConfiguration"`
}

