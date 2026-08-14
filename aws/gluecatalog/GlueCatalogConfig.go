// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#name GlueCatalog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#allow_full_table_external_data_access GlueCatalog#allow_full_table_external_data_access}.
	AllowFullTableExternalDataAccess *string `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// catalog_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#catalog_properties GlueCatalog#catalog_properties}
	CatalogProperties interface{} `field:"optional" json:"catalogProperties" yaml:"catalogProperties"`
	// create_database_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#create_database_default_permissions GlueCatalog#create_database_default_permissions}
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// create_table_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#create_table_default_permissions GlueCatalog#create_table_default_permissions}
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#description GlueCatalog#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// federated_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#federated_catalog GlueCatalog#federated_catalog}
	FederatedCatalog interface{} `field:"optional" json:"federatedCatalog" yaml:"federatedCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#overwrite_child_resource_permissions_with_default GlueCatalog#overwrite_child_resource_permissions_with_default}.
	OverwriteChildResourcePermissionsWithDefault *string `field:"optional" json:"overwriteChildResourcePermissionsWithDefault" yaml:"overwriteChildResourcePermissionsWithDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#parameters GlueCatalog#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#region GlueCatalog#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#tags GlueCatalog#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// target_redshift_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#target_redshift_catalog GlueCatalog#target_redshift_catalog}
	TargetRedshiftCatalog interface{} `field:"optional" json:"targetRedshiftCatalog" yaml:"targetRedshiftCatalog"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/glue_catalog#timeouts GlueCatalog#timeouts}
	Timeouts *GlueCatalogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

