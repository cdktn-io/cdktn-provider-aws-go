// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawskmssecrets


type EphemeralAwsKmsSecretsSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#name EphemeralAwsKmsSecrets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#payload EphemeralAwsKmsSecrets#payload}.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#context EphemeralAwsKmsSecrets#context}.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#encryption_algorithm EphemeralAwsKmsSecrets#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#grant_tokens EphemeralAwsKmsSecrets#grant_tokens}.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#key_id EphemeralAwsKmsSecrets#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

