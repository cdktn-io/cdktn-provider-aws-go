// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/glue_catalog_table#fields GlueCatalogTable#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/glue_catalog_table#spec_id GlueCatalogTable#spec_id}.
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

