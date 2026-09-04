// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationDeletionProtectionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagent_data_source#deletion_protection_status BedrockagentDataSource#deletion_protection_status}.
	DeletionProtectionStatus *string `field:"required" json:"deletionProtectionStatus" yaml:"deletionProtectionStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagent_data_source#deletion_protection_threshold BedrockagentDataSource#deletion_protection_threshold}.
	DeletionProtectionThreshold *float64 `field:"optional" json:"deletionProtectionThreshold" yaml:"deletionProtectionThreshold"`
}

