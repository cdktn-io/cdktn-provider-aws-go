// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrderFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog_table#direction GlueCatalogTable#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog_table#null_order GlueCatalogTable#null_order}.
	NullOrder *string `field:"required" json:"nullOrder" yaml:"nullOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog_table#source_id GlueCatalogTable#source_id}.
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog_table#transform GlueCatalogTable#transform}.
	Transform *string `field:"required" json:"transform" yaml:"transform"`
}

