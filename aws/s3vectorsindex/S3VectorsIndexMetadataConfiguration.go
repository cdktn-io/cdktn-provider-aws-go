// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsindex


type S3VectorsIndexMetadataConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/s3vectors_index#non_filterable_metadata_keys S3VectorsIndex#non_filterable_metadata_keys}.
	NonFilterableMetadataKeys *[]*string `field:"required" json:"nonFilterableMetadataKeys" yaml:"nonFilterableMetadataKeys"`
}

