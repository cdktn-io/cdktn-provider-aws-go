// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogCatalogPropertiesIcebergOptimizationProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/glue_catalog#compaction GlueCatalog#compaction}.
	Compaction *map[string]*string `field:"optional" json:"compaction" yaml:"compaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/glue_catalog#orphan_file_deletion GlueCatalog#orphan_file_deletion}.
	OrphanFileDeletion *map[string]*string `field:"optional" json:"orphanFileDeletion" yaml:"orphanFileDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/glue_catalog#retention GlueCatalog#retention}.
	Retention *map[string]*string `field:"optional" json:"retention" yaml:"retention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/glue_catalog#role_arn GlueCatalog#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

