// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigMicrosoftOauth2ProviderConfigClientSecretConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#json_key BedrockagentcoreOauth2CredentialProvider#json_key}.
	JsonKey *string `field:"required" json:"jsonKey" yaml:"jsonKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#secret_id BedrockagentcoreOauth2CredentialProvider#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

