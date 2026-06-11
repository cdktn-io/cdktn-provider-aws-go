// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/glue_catalog_table#metadata_operation GlueCatalogTable#metadata_operation}.
	MetadataOperation *string `field:"required" json:"metadataOperation" yaml:"metadataOperation"`
	// iceberg_table_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/glue_catalog_table#iceberg_table_input GlueCatalogTable#iceberg_table_input}
	IcebergTableInput *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput `field:"optional" json:"icebergTableInput" yaml:"icebergTableInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/glue_catalog_table#version GlueCatalogTable#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

