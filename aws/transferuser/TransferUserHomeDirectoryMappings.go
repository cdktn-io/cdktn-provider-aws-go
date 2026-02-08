// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferuser


type TransferUserHomeDirectoryMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/transfer_user#entry TransferUser#entry}.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/transfer_user#target TransferUser#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
}

