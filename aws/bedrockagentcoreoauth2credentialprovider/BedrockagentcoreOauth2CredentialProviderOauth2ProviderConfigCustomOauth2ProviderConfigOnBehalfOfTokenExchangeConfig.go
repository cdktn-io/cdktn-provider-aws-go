// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigOnBehalfOfTokenExchangeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#grant_type BedrockagentcoreOauth2CredentialProvider#grant_type}.
	GrantType *string `field:"required" json:"grantType" yaml:"grantType"`
	// token_exchange_grant_type_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_oauth2_credential_provider#token_exchange_grant_type_config BedrockagentcoreOauth2CredentialProvider#token_exchange_grant_type_config}
	TokenExchangeGrantTypeConfig interface{} `field:"optional" json:"tokenExchangeGrantTypeConfig" yaml:"tokenExchangeGrantTypeConfig"`
}

