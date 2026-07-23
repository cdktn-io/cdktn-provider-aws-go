// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionAuthenticationConfigurationOauth2Properties struct {
	// authorization_code_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/glue_connection#authorization_code_properties GlueConnection#authorization_code_properties}
	AuthorizationCodeProperties *GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// oauth2_client_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/glue_connection#oauth2_client_application GlueConnection#oauth2_client_application}
	Oauth2ClientApplication *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication `field:"optional" json:"oauth2ClientApplication" yaml:"oauth2ClientApplication"`
	// oauth2_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/glue_connection#oauth2_credentials GlueConnection#oauth2_credentials}
	Oauth2Credentials *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials `field:"optional" json:"oauth2Credentials" yaml:"oauth2Credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/glue_connection#oauth2_grant_type GlueConnection#oauth2_grant_type}.
	Oauth2GrantType *string `field:"optional" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/glue_connection#token_url GlueConnection#token_url}.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/glue_connection#token_url_parameters_map GlueConnection#token_url_parameters_map}.
	TokenUrlParametersMap *map[string]*string `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

