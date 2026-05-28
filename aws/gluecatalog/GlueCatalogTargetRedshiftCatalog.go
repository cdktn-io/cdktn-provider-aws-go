// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogTargetRedshiftCatalog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/glue_catalog#catalog_arn GlueCatalog#catalog_arn}.
	CatalogArn *string `field:"required" json:"catalogArn" yaml:"catalogArn"`
}

