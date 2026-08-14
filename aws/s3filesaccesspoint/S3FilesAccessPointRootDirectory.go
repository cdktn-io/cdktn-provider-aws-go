// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesaccesspoint


type S3FilesAccessPointRootDirectory struct {
	// creation_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/s3files_access_point#creation_permissions S3FilesAccessPoint#creation_permissions}
	CreationPermissions interface{} `field:"optional" json:"creationPermissions" yaml:"creationPermissions"`
	// Root directory path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/s3files_access_point#path S3FilesAccessPoint#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

