// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_catalog_table#location GlueCatalogTable#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_catalog_table#schema GlueCatalogTable#schema}
	Schema *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema `field:"required" json:"schema" yaml:"schema"`
	// partition_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_catalog_table#partition_spec GlueCatalogTable#partition_spec}
	PartitionSpec *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_catalog_table#properties GlueCatalogTable#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// sort_order block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_catalog_table#sort_order GlueCatalogTable#sort_order}
	SortOrder *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

