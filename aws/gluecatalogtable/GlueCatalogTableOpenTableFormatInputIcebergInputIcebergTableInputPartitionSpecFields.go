// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/glue_catalog_table#name GlueCatalogTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/glue_catalog_table#source_id GlueCatalogTable#source_id}.
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/glue_catalog_table#transform GlueCatalogTable#transform}.
	Transform *string `field:"required" json:"transform" yaml:"transform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/glue_catalog_table#field_id GlueCatalogTable#field_id}.
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
}

