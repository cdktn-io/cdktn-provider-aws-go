// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogCatalogPropertiesDataLakeAccessProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_catalog#catalog_type GlueCatalog#catalog_type}.
	CatalogType *string `field:"optional" json:"catalogType" yaml:"catalogType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_catalog#data_lake_access GlueCatalog#data_lake_access}.
	DataLakeAccess interface{} `field:"optional" json:"dataLakeAccess" yaml:"dataLakeAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_catalog#data_transfer_role GlueCatalog#data_transfer_role}.
	DataTransferRole *string `field:"optional" json:"dataTransferRole" yaml:"dataTransferRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_catalog#kms_key GlueCatalog#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

