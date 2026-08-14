// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfigurationCompactionConfigurationIcebergConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog_table_optimizer#delete_file_threshold GlueCatalogTableOptimizer#delete_file_threshold}.
	DeleteFileThreshold *float64 `field:"optional" json:"deleteFileThreshold" yaml:"deleteFileThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog_table_optimizer#min_input_files GlueCatalogTableOptimizer#min_input_files}.
	MinInputFiles *float64 `field:"optional" json:"minInputFiles" yaml:"minInputFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog_table_optimizer#strategy GlueCatalogTableOptimizer#strategy}.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

