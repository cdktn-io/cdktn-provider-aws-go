// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetCredentialProviderConfigurationOauth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#provider_arn BedrockagentcoreGatewayTarget#provider_arn}.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#scopes BedrockagentcoreGatewayTarget#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#custom_parameters BedrockagentcoreGatewayTarget#custom_parameters}.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
	// The URL where the end user's browser is redirected after obtaining the authorization code. Required when grant_type is AUTHORIZATION_CODE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#default_return_url BedrockagentcoreGatewayTarget#default_return_url}
	DefaultReturnUrl *string `field:"optional" json:"defaultReturnUrl" yaml:"defaultReturnUrl"`
	// The OAuth grant type. Valid values are AUTHORIZATION_CODE and CLIENT_CREDENTIALS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway_target#grant_type BedrockagentcoreGatewayTarget#grant_type}
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
}

