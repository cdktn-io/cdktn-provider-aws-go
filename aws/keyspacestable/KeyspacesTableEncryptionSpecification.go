// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTableEncryptionSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/keyspaces_table#kms_key_identifier KeyspacesTable#kms_key_identifier}.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/keyspaces_table#type KeyspacesTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

