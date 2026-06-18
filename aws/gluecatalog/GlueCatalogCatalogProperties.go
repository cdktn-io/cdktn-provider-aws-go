// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogCatalogProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog#custom_properties GlueCatalog#custom_properties}.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
	// data_lake_access_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog#data_lake_access_properties GlueCatalog#data_lake_access_properties}
	DataLakeAccessProperties interface{} `field:"optional" json:"dataLakeAccessProperties" yaml:"dataLakeAccessProperties"`
	// iceberg_optimization_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/glue_catalog#iceberg_optimization_properties GlueCatalog#iceberg_optimization_properties}
	IcebergOptimizationProperties interface{} `field:"optional" json:"icebergOptimizationProperties" yaml:"icebergOptimizationProperties"`
}

