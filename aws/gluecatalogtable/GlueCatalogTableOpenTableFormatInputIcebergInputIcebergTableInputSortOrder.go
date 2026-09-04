// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_catalog_table#fields GlueCatalogTable#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_catalog_table#order_id GlueCatalogTable#order_id}.
	OrderId *float64 `field:"required" json:"orderId" yaml:"orderId"`
}

