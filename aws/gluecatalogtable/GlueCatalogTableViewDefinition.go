// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableViewDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#definer GlueCatalogTable#definer}.
	Definer *string `field:"optional" json:"definer" yaml:"definer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#is_protected GlueCatalogTable#is_protected}.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#last_refresh_type GlueCatalogTable#last_refresh_type}.
	LastRefreshType *string `field:"optional" json:"lastRefreshType" yaml:"lastRefreshType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#refresh_seconds GlueCatalogTable#refresh_seconds}.
	RefreshSeconds *float64 `field:"optional" json:"refreshSeconds" yaml:"refreshSeconds"`
	// representations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#representations GlueCatalogTable#representations}
	Representations interface{} `field:"optional" json:"representations" yaml:"representations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#sub_objects GlueCatalogTable#sub_objects}.
	SubObjects *[]*string `field:"optional" json:"subObjects" yaml:"subObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#sub_object_version_ids GlueCatalogTable#sub_object_version_ids}.
	SubObjectVersionIds *[]*float64 `field:"optional" json:"subObjectVersionIds" yaml:"subObjectVersionIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#view_version_id GlueCatalogTable#view_version_id}.
	ViewVersionId *float64 `field:"optional" json:"viewVersionId" yaml:"viewVersionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/glue_catalog_table#view_version_token GlueCatalogTable#view_version_token}.
	ViewVersionToken *string `field:"optional" json:"viewVersionToken" yaml:"viewVersionToken"`
}

