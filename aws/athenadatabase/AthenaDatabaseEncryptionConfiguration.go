// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenadatabase


type AthenaDatabaseEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/athena_database#encryption_option AthenaDatabase#encryption_option}.
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/athena_database#kms_key AthenaDatabase#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

