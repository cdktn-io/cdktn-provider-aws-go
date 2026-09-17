// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigPrivateKeyJwtConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#additional_header_claims BedrockagentcoreOauth2CredentialProvider#additional_header_claims}.
	AdditionalHeaderClaims *map[string]*string `field:"optional" json:"additionalHeaderClaims" yaml:"additionalHeaderClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#additional_payload_claims BedrockagentcoreOauth2CredentialProvider#additional_payload_claims}.
	AdditionalPayloadClaims *map[string]*string `field:"optional" json:"additionalPayloadClaims" yaml:"additionalPayloadClaims"`
	// private_key_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#private_key_source BedrockagentcoreOauth2CredentialProvider#private_key_source}
	PrivateKeySource interface{} `field:"optional" json:"privateKeySource" yaml:"privateKeySource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#signing_algorithm BedrockagentcoreOauth2CredentialProvider#signing_algorithm}.
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
}

