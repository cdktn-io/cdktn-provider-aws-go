// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationDestinationTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/msk_channel#destination_database_name MskChannel#destination_database_name}.
	DestinationDatabaseName *string `field:"optional" json:"destinationDatabaseName" yaml:"destinationDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/msk_channel#destination_table_name MskChannel#destination_table_name}.
	DestinationTableName *string `field:"optional" json:"destinationTableName" yaml:"destinationTableName"`
	// partition_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/msk_channel#partition_spec MskChannel#partition_spec}
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
}

