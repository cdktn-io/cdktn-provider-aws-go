// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_connection#authentication_type GlueConnection#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// basic_authentication_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_connection#basic_authentication_credentials GlueConnection#basic_authentication_credentials}
	BasicAuthenticationCredentials *GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentials `field:"optional" json:"basicAuthenticationCredentials" yaml:"basicAuthenticationCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_connection#custom_authentication_credentials GlueConnection#custom_authentication_credentials}.
	CustomAuthenticationCredentials *map[string]*string `field:"optional" json:"customAuthenticationCredentials" yaml:"customAuthenticationCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_connection#kms_key_arn GlueConnection#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// oauth2_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_connection#oauth2_properties GlueConnection#oauth2_properties}
	Oauth2Properties *GlueConnectionAuthenticationConfigurationOauth2Properties `field:"optional" json:"oauth2Properties" yaml:"oauth2Properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/glue_connection#secret_arn GlueConnection#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

