// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMetadataIceberg struct {
	// A map of configuration properties for the Iceberg table, for example `write.distribution-mode` and `write.sort-order`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/s3tables_table#properties S3TablesTable#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/s3tables_table#schema S3TablesTable#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

