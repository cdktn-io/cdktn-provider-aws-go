// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogFederatedCatalog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/glue_catalog#connection_name GlueCatalog#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/glue_catalog#connection_type GlueCatalog#connection_type}.
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/glue_catalog#identifier GlueCatalog#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

