// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyDayOfMonth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/bedrockagent_data_source#day_number BedrockagentDataSource#day_number}.
	DayNumber *float64 `field:"optional" json:"dayNumber" yaml:"dayNumber"`
	// last_day_of_month block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/bedrockagent_data_source#last_day_of_month BedrockagentDataSource#last_day_of_month}
	LastDayOfMonth interface{} `field:"optional" json:"lastDayOfMonth" yaml:"lastDayOfMonth"`
}

