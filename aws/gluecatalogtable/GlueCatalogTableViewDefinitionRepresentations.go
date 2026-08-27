// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableViewDefinitionRepresentations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#dialect GlueCatalogTable#dialect}.
	Dialect *string `field:"optional" json:"dialect" yaml:"dialect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#dialect_version GlueCatalogTable#dialect_version}.
	DialectVersion *string `field:"optional" json:"dialectVersion" yaml:"dialectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#validation_connection GlueCatalogTable#validation_connection}.
	ValidationConnection *string `field:"optional" json:"validationConnection" yaml:"validationConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_expanded_text GlueCatalogTable#view_expanded_text}.
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_original_text GlueCatalogTable#view_original_text}.
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

