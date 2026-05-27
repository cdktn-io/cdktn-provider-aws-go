// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfigurationOrphanFileDeletionConfiguration struct {
	// iceberg_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/glue_catalog_table_optimizer#iceberg_configuration GlueCatalogTableOptimizer#iceberg_configuration}
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

