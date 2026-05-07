// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/glue_connection#authorization_code GlueConnection#authorization_code}.
	AuthorizationCode *string `field:"required" json:"authorizationCode" yaml:"authorizationCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/glue_connection#redirect_uri GlueConnection#redirect_uri}.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

