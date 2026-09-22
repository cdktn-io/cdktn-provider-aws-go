// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOauth2CredentialProviderOauth2ProviderConfigCustomOauth2ProviderConfigOnBehalfOfTokenExchangeConfigTokenExchangeGrantTypeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/bedrockagentcore_oauth2_credential_provider#actor_token_content BedrockagentcoreOauth2CredentialProvider#actor_token_content}.
	ActorTokenContent *string `field:"required" json:"actorTokenContent" yaml:"actorTokenContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/bedrockagentcore_oauth2_credential_provider#actor_token_scopes BedrockagentcoreOauth2CredentialProvider#actor_token_scopes}.
	ActorTokenScopes *[]*string `field:"optional" json:"actorTokenScopes" yaml:"actorTokenScopes"`
}

